package serverquery

import "fmt"

type ChannelID int

type Channel struct {
	HostingServer VirtualServer
	stale         bool
	ID            ChannelID `sq:"cid"`
	PID           int       `sq:"pid"`
	Order         int       `sq:"channel_order"`
	Name          string    `sq:"channel_name"`
	ClientsOnline int       `sq:"total_clients"`
	MaxClients    int       `sq:"channel_maxclients"`
	Codec         int       `sq:"channel_codec"`
	CodecQuality  int       `sq:"channel_codec_quality"`
	LatencyFactor int       `sq:"channel_codec_latency_factor"`
	Unencrypted   int       `sq:"channel_codec_is_unencrypted"`
	Permanent     int       `sq:"channel_flag_permanent"`
	SemiPermanent int       `sq:"channel_flag_semi_permanent"`
	Default       int       `sq:"channel_flag_default"`
	Password      int       `sq:"channel_flag_password"`
}

// VServerInventory can return all known VServer
type VServerInventory interface {
	All() []VirtualServer
}

type ChannelView struct {
	e        Executor
	channels map[ChannelID]Channel
	vServer  VServerInventory
}

func NewChannelView(e Executor, inventory VServerInventory) *ChannelView {
	return &ChannelView{
		e:        e,
		channels: make(map[ChannelID]Channel),
		vServer:  inventory,
	}
}

// Refresh refreshes the internal representation of the ChannelView. It changes into all virtual server
// known by the vServer inventory
func (c *ChannelView) Refresh() error {
	c.markAllStale()
	defer c.deleteAllStale()
	for _, vServer := range c.vServer.All() {
		err := c.updateAllOnVServer(vServer)
		if err != nil {
			return fmt.Errorf("failed to update metrics on vServer %s: %w", vServer.Name, err)
		}
	}
	return nil
}

// All returns all known Channels
func (c *ChannelView) All() []Channel {
	ret := make([]Channel, 0, len(c.channels))
	for _, ch := range c.channels {
		ret = append(ret, ch)
	}
	return ret
}

// markAllStale marks all channels stale. They are set during scrape. If the aren't set, they will be deleted
// by deleteAllStale
func (c *ChannelView) markAllStale() {
	for id, channel := range c.channels {
		channel.stale = true
		c.channels[id] = channel
	}
}

// deleteAllStale deletes all stale channels
func (c *ChannelView) deleteAllStale() {
	for id, channel := range c.channels {
		if channel.stale {
			delete(c.channels, id)
		}
	}
}

// updateAllOnVServer update all channels on the given virtual server
func (c *ChannelView) updateAllOnVServer(vServer VirtualServer) error {
	_, err := c.e.Exec(fmt.Sprintf("use %d", vServer.ID))
	if err != nil {
		return fmt.Errorf("failed to use virtual server %d: %w", vServer.ID, err)
	}
	res, err := c.e.Exec("channellist")
	if err != nil {
		return fmt.Errorf("failed to list channels: %w", err)
	}
	for _, r := range res {
		for _, i := range r.Items {
			var ch Channel
			if err := i.ReadInto(&ch); err != nil {
				return fmt.Errorf("failed to parse channel from response: %w", err)
			}
			if err := c.getDetails(&ch); err != nil {
				return fmt.Errorf("failed to parse details for channel %d: %w", ch.ID, err)
			}
			ch.HostingServer = vServer
			c.channels[ch.ID] = ch
		}
	}
	return nil
}

// getDetails populates the given channels with details
func (c *ChannelView) getDetails(ch *Channel) error {
	res, err := c.e.Exec(fmt.Sprintf("channelinfo cid=%d", ch.ID))
	if err != nil {
		return fmt.Errorf("failed to run channelinfo command: %w", err)
	}
	for _, r := range res {
		if len(r.Items) != 1 {
			return fmt.Errorf("expected exactly one channelinfo response got %d", len(r.Items))
		}
		if err = r.Items[0].ReadInto(ch); err != nil {
			return fmt.Errorf("failed to parse channel response: %w", err)
		}
		return nil
	}
	return fmt.Errorf("reached unreachable code")
}
