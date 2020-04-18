package serverquery

import (
	"errors"
	"fmt"
	"syscall"
	"time"

	"github.com/multiplay/go-ts3"
)

type Client struct {
	ts3Client          *ts3.Client
	limiter            *time.Ticker
	user               string
	password           string
	remote             string
	respectLimits      bool
	serverQueryOptions []options
	metrics            *ClientMetrics
}

type options func() func(client *ts3.Client) error

// NewClient instantiate a new client to the given remote. If the function returns successfully the client is already
// logged in.
func NewClient(remote, user, password string, ignoreLimits bool, serverQueryOptions ...options) (*Client, error) {
	c := &Client{
		user:               user,
		password:           password,
		remote:             remote,
		respectLimits:      !ignoreLimits,
		serverQueryOptions: serverQueryOptions,
		metrics:            &ClientMetrics{},
	}
	return c, c.reconnect()
}

func (c *Client) reconnect() error {
	if err := c.login(c.remote, c.serverQueryOptions...); err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	if c.respectLimits {
		if err := c.setupLimiter(); err != nil {
			return fmt.Errorf("failed to setup limiter: %w", err)
		}
	}
	return nil
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

func (c *Client) setupLimiter() error {
	type instanceInfo struct {
		TimeWindow float64 `sq:"serverinstance_serverquery_flood_time"`
		Commands   float64 `sq:"serverinstance_serverquery_flood_commands"`
	}
	res, err := c.Exec("instanceinfo")
	if err != nil {
		return fmt.Errorf("failed to execute instanceinfo command: %w", err)
	}
	if len(res) != 1 {
		return fmt.Errorf("expected exactly one response got: %d", len(res))
	}
	if len(res[0].Items) != 1 {
		return fmt.Errorf("expected exactly one item got %d", len(res[0].Items))
	}
	var ii instanceInfo
	if err := res[0].Items[0].ReadInto(&ii); err != nil {
		return fmt.Errorf("failed to parse answer into a instance info")
	}
	// add 10 % buffer to not run into the flood limit
	// multiply by 1000 to get milliseconds. Most probably the flood limit interval is several hundred milliseconds
	interval := time.Duration((ii.TimeWindow/ii.Commands)*1100) * time.Millisecond
	c.limiter = time.NewTicker(interval)
	return nil
}

func (c *Client) Metrics() *ClientMetrics {
	return c.metrics
}

func (c *Client) Exec(cmd string) ([]Result, error) {
	if c.limiter != nil {
		<-c.limiter.C
	}
	if c.ts3Client == nil {
		err := c.reconnect()
		if err != nil {
			return nil, fmt.Errorf("failed to reconnect: %w", err)
		}
	}
	raw, err := c.ts3Client.Exec(cmd)
	if err != nil {
		c.metrics.CountFailure()
		// If pipe broke, reconnect on next execution
		if errors.Is(err, syscall.EPIPE) {
			c.ts3Client = nil
		}
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	ret := make([]Result, 0, len(raw))
	for _, r := range raw {
		p, err := Parse(r)
		if err != nil {
			c.metrics.CountFailure()
			return nil, fmt.Errorf("failed to parse answer: %w", err)
		}
		ret = append(ret, p)
	}
	c.metrics.CountSuccess()
	return ret, nil
}
