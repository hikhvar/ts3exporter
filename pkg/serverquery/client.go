package serverquery

import (
	"fmt"

	"github.com/multiplay/go-ts3"
)

type Client struct {
	ts3Client       *ts3.Client
	user            string
	password        string
	virtualServerID int
}

type options func() func(client *ts3.Client) error

// NewClient instantiate a new client to the given remote. If the function returns successfully the client is already
// logged in.
func NewClient(remote, user, password string, serverQueryOptions ...options) (*Client, error) {
	c := &Client{
		user:     user,
		password: password,
	}
	err := c.login(remote, serverQueryOptions...)
	return c, err
}

// login connects to the given remote with the given options.
func (c *Client) login(remote string, serverQueryOptions ...options) error {
	upstreamOptions := make([]func(client *ts3.Client) error, 0, len(serverQueryOptions))
	for _, opt := range serverQueryOptions {
		upstreamOptions = append(upstreamOptions, opt())
	}
	t, err := ts3.NewClient(remote, upstreamOptions...)
	if err != nil {
		return fmt.Errorf("failed to connect to TS3 serverquery endpoint: %w", err)
	}
	c.ts3Client = t

	err = t.Login(c.user, c.password)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	return nil
}

func (c *Client) Exec(cmd string) ([]Result, error) {
	raw, err := c.ts3Client.Exec(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	ret := make([]Result, 0, len(raw))
	for _, r := range raw {
		p, err := Parse(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse answer: %w", err)
		}
		ret = append(ret, p)
	}
	return ret, nil
}
