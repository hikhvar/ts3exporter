package serverquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const channelInput = `pid=10 channel_name=Default\sChannel channel_topic=Default\sChannel\shas\sno\stopic channel_description=This\sis\sthe\sdefault\schannel channel_password channel_codec=4 channel_codec_quality=6 channel_maxclients=-1 channel_maxfamilyclients=-1 channel_order=11 channel_flag_permanent=1 channel_flag_semi_permanent=0 channel_flag_default=1 channel_flag_password=0 channel_codec_latency_factor=1 channel_codec_is_unencrypted=1 channel_security_salt channel_delete_delay=0 channel_unique_identifier=d1eb7624-664a-4809-9b7e-84596d937a6d channel_flag_maxclients_unlimited=1 channel_flag_maxfamilyclients_unlimited=1 channel_flag_maxfamilyclients_inherited=0 channel_filepath=files\/virtualserver_1\/channel_1 channel_needed_talk_power=0 channel_forced_silence=0 channel_name_phonetic channel_icon_id=0 channel_banner_gfx_url channel_banner_mode=0 seconds_empty=-1`

func TestChannelStructTags(t *testing.T) {
	res, err := Parse(channelInput)
	require.Nil(t, err)
	require.Len(t, res.Items, 1)
	ch := Channel{
		ClientsOnline: 5,
		ID:            1,
	}
	err = res.Items[0].ReadInto(&ch)
	require.Nil(t, err)
	expected := Channel{
		HostingServer: VirtualServer{},
		ID:            1,
		PID:           10,
		Order:         11,
		Name:          "Default Channel",
		ClientsOnline: 5,
		MaxClients:    -1,
		Codec:         4,
		CodecQuality:  6,
		LatencyFactor: 1,
		Unencrypted:   1,
		Permanent:     1,
		SemiPermanent: 0,
		Default:       1,
		Password:      0,
	}
	assert.Equal(t, expected, ch)
}
