package serverquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const input = `virtualserver_unique_identifier=qnCQiIJ8Z0Bdc\/N8MqIUyjZDCMw= virtualserver_name=Gute\sStube virtualserver_welcomemessage=Welcome\sto\sTeamSpeak,\scheck\s[URL]www.teamspeak.com[\/URL]\sfor\slatest\sinformation virtualserver_platform=Linux virtualserver_version=3.11.0\s[Build:\s1578903157] virtualserver_maxclients=32 virtualserver_password=\/YI\/s3mS9GVawcLNudHDaI9p5X4= virtualserver_clientsonline=3 virtualserver_channelsonline=6 virtualserver_created=1569786227 virtualserver_uptime=11088 virtualserver_codec_encryption_mode=0 virtualserver_hostmessage virtualserver_hostmessage_mode=0 virtualserver_filebase=files\/virtualserver_1 virtualserver_default_server_group=8 virtualserver_default_channel_group=8 virtualserver_flag_password=1 virtualserver_default_channel_admin_group=5 virtualserver_max_download_total_bandwidth=18446744077709551615 virtualserver_max_upload_total_bandwidth=18446744073709551615 virtualserver_hostbanner_url virtualserver_hostbanner_gfx_url virtualserver_hostbanner_gfx_interval=0 virtualserver_complain_autoban_count=5 virtualserver_complain_autoban_time=1200 virtualserver_complain_remove_time=3600 virtualserver_min_clients_in_channel_before_forced_silence=100 virtualserver_priority_speaker_dimm_modificator=-18.0000 virtualserver_id=1 virtualserver_antiflood_points_tick_reduce=5 virtualserver_antiflood_points_needed_command_block=150 virtualserver_antiflood_points_needed_ip_block=250 virtualserver_client_connections=1 virtualserver_query_client_connections=6 virtualserver_hostbutton_tooltip virtualserver_hostbutton_url virtualserver_hostbutton_gfx_url virtualserver_queryclientsonline=1 virtualserver_download_quota=18446744073709551615 virtualserver_upload_quota=18446744073709551615 virtualserver_month_bytes_downloaded=0 virtualserver_month_bytes_uploaded=0 virtualserver_total_bytes_downloaded=0 virtualserver_total_bytes_uploaded=0 virtualserver_port=9987 virtualserver_autostart=1 virtualserver_machine_id virtualserver_needed_identity_security_level=8 virtualserver_log_client=0 virtualserver_log_query=0 virtualserver_log_channel=0 virtualserver_log_permissions=1 virtualserver_log_server=0 virtualserver_log_filetransfer=0 virtualserver_min_client_version=1560850141 virtualserver_name_phonetic virtualserver_icon_id=0 virtualserver_reserved_slots=0 virtualserver_total_packetloss_speech=0.0000 virtualserver_total_packetloss_keepalive=0.0000 virtualserver_total_packetloss_control=0.0000 virtualserver_total_packetloss_total=0.0000 virtualserver_total_ping=0.0000 virtualserver_ip=0.0.0.0,\s:: virtualserver_weblist_enabled=1 virtualserver_ask_for_privilegekey=0 virtualserver_hostbanner_mode=0 virtualserver_channel_temp_delete_delay_default=0 virtualserver_min_android_version=1559834030 virtualserver_min_ios_version=1559144369 virtualserver_nickname virtualserver_antiflood_points_needed_plugin_block=0 virtualserver_status=online connection_filetransfer_bandwidth_sent=0 connection_filetransfer_bandwidth_received=0 connection_filetransfer_bytes_sent_total=0 connection_filetransfer_bytes_received_total=0 connection_packets_sent_speech=0 connection_bytes_sent_speech=0 connection_packets_received_speech=34 connection_bytes_received_speech=3501 connection_packets_sent_keepalive=367 connection_bytes_sent_keepalive=15047 connection_packets_received_keepalive=367 connection_bytes_received_keepalive=15413 connection_packets_sent_control=41 connection_bytes_sent_control=8189 connection_packets_received_control=44 connection_bytes_received_control=4182 connection_packets_sent_total=408 connection_bytes_sent_total=23236 connection_packets_received_total=445 connection_bytes_received_total=23096 connection_bandwidth_sent_last_second_total=0 connection_bandwidth_sent_last_minute_total=0 connection_bandwidth_received_last_second_total=0 connection_bandwidth_received_last_minute_total=0`

func TestVirtualServerStructTags(t *testing.T) {
	res, err := Parse(input)
	require.Nil(t, err)
	var v VirtualServer
	err = res.Items[0].ReadInto(&v)
	require.Nil(t, err)
	expected := VirtualServer{
		ID:                             1,
		Port:                           9987,
		Name:                           "Gute Stube",
		Status:                         "online",
		ClientsOnline:                  3,
		QueryClientsOnline:             1,
		MaxClients:                     32,
		Uptime:                         11088,
		ChannelsOnline:                 6,
		MaxDownloadTotalBandwidth:      1.844674407770955e+19,
		MaxUploadTotalBandwidth:        1.8446744073709552e+19,
		ClientsConnections:             1,
		QueryClientsConnections:        1,
		FileTransferBytesSentTotal:     0,
		FileTransferBytesReceivedTotal: 0,
		ControlBytesSendTotal:          8189,
		ControlBytesReceivedTotal:      4182,
		SpeechBytesSendTotal:           0,
		SpeechBytesReceivedTotal:       3501,
		KeepAliveBytesSendTotal:        15047,
		KeepAliveBytesReceivedTotal:    15413,
		BytesSendTotal:                 23236,
		BytesReceivedTotal:             23096,
	}
	assert.Equal(t, expected, v)
}
