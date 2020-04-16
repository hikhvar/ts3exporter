package serverquery

import (
	"fmt"
)

type VirtualServerID int

type VirtualServer struct {
	ID                        VirtualServerID `sq:"virtualserver_id"`
	Port                      int             `sq:"virtualserver_port"`
	Name                      string          `sq:"virtualserver_name"`
	Status                    string          `sq:"virtualserver_status"`
	ClientsOnline             int             `sq:"virtualserver_clientsonline"`
	QueryClientsOnline        int             `sq:"virtualserver_queryclientsonline"`
	MaxClients                int             `sq:"virtualserver_maxclients"`
	Uptime                    int             `sq:"virtualserver_uptime"`
	ChannelsOnline            int             `sq:"virtualserver_channelsonline"`
	MaxDownloadTotalBandwidth float64         `sq:"virtualserver_max_download_total_bandwidth"`
	MaxUploadTotalBandwidth   float64         `sq:"virtualserver_max_upload_total_bandwidth"`
	ClientsConnections        int             `sq:"virtualserver_client_connections"`
	QueryClientsConnections   int             `sq:"virtualserver_queryclientsonline"`

	FileTransferBytesSentTotal     int `sq:"connection_filetransfer_bytes_sent_total"`
	FileTransferBytesReceivedTotal int `sq:"connection_filetransfer_bytes_received_total"`

	ControlBytesSendTotal     int `sq:"connection_bytes_sent_control"`
	ControlBytesReceivedTotal int `sq:"connection_bytes_received_control"`

	SpeechBytesSendTotal     int `sq:"connection_bytes_sent_speech"`
	SpeechBytesReceivedTotal int `sq:"connection_bytes_received_speech"`

	KeepAliveBytesSendTotal     int `sq:"connection_bytes_sent_keepalive"`
	KeepAliveBytesReceivedTotal int `sq:"connection_bytes_received_keepalive"`

	BytesSendTotal     int `sq:"connection_bytes_sent_total"`
	BytesReceivedTotal int `sq:"connection_bytes_received_total"`
}

type VirtualServerView struct {
	e       Executor
	vServer map[VirtualServerID]VirtualServer
}

func NewVirtualServer(e Executor) *VirtualServerView {
	return &VirtualServerView{
		e:       e,
		vServer: make(map[VirtualServerID]VirtualServer),
	}
}

// Refresh refreshes the internal representation of the VirtualServerView
func (v *VirtualServerView) Refresh() error {
	v.vServer = make(map[VirtualServerID]VirtualServer, len(v.vServer))
	res, err := v.e.Exec("serverlist")
	if err != nil {
		return fmt.Errorf("failed to list v servers: %w", err)
	}
	for _, r := range res {
		var vs VirtualServer
		for _, i := range r.Items {
			if err = i.ReadInto(&vs); err != nil {
				return fmt.Errorf("failed to read virtual server from return: %w", err)
			}
			vs, err = v.getDetails(vs.ID)
			if err != nil {
				return fmt.Errorf("failed to fetch details: %w", err)
			}
			v.vServer[vs.ID] = vs
		}
	}
	return nil
}

// getDetails uses the serverinfo serverquery command to get the details of the given virtualserver
func (v *VirtualServerView) getDetails(vServerID VirtualServerID) (VirtualServer, error) {
	_, err := v.e.Exec(fmt.Sprintf("use %d", vServerID))
	if err != nil {
		return VirtualServer{}, fmt.Errorf("failed to use virtual server %d: %w", vServerID, err)
	}
	res, err := v.e.Exec("serverinfo")
	if err != nil {
		return VirtualServer{}, fmt.Errorf("failed to fetch serverinfo for virtual server %d: %w", vServerID, err)
	}
	for _, r := range res {
		var vs VirtualServer
		for _, i := range r.Items {
			if err = i.ReadInto(&vs); err != nil {
				return VirtualServer{}, fmt.Errorf("failed to read virtual server from return: %w", err)
			}
			return vs, nil
		}
	}
	return VirtualServer{}, fmt.Errorf("reached unreachable code")
}

// All returns all known virtual server
func (v *VirtualServerView) All() []VirtualServer {
	ret := make([]VirtualServer, 0, len(v.vServer))
	for _, vs := range v.vServer {
		ret = append(ret, vs)
	}
	return ret
}
