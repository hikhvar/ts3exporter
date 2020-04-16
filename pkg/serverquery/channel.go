package serverquery

import "fmt"

type ChannelID int

type Channel struct {
	HostingServer VirtualServer
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

type ChannelView struct {
	e                Executor
	channels         map[ChannelID]Channel
	vServerInventory *VirtualServerView
}

func NewChannelView(e Executor) *ChannelView {
	return &ChannelView{
		e:                e,
		channels:         make(map[ChannelID]Channel),
		vServerInventory: NewVirtualServer(e),
	}
}

// Refresh refreshes the internal representation of the ChannelView. It changes into all virtual server
// known by the vServer inventory.
func (c *ChannelView) Refresh() error {
	if err := c.vServerInventory.Refresh(); err != nil {
		return fmt.Errorf("failed to update vserver inventory: %w", err)
	}
	for _, vServer := range c.vServerInventory.All() {
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
	if len(res) < 1 {
		return fmt.Errorf("expected at least one response line")
	}
	if len(res[0].Items) != 1 {
		return fmt.Errorf("expected exactly one channelinfo response got %d", len(res[0].Items))
	}
	if err = res[0].Items[0].ReadInto(ch); err != nil {
		return fmt.Errorf("failed to parse channel response: %w", err)
	}
	return nil

}
