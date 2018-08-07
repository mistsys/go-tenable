package tenable

import (
	"context"
	"fmt"
)

type ScannersService service

// represents an instance of a scanner
type Scanner struct {
	CreationDate         int    `json:"creation_date"`
	Distro               string `json:"distro,omitempty"`
	EngineVersion        string `json:"engine_version,omitempty"`
	Group                bool   `json:"group"`
	ID                   int    `json:"id"`
	Key                  string `json:"key"`
	LastConnect          int    `json:"last_connect"`
	LastModificationDate int    `json:"last_modification_date"`
	Linked               int    `json:"linked"`
	LoadedPluginSet      string `json:"loaded_plugin_set,omitempty"`
	Name                 string `json:"name"`
	EnvironmentName      string `json:"environment_name"`
	NumHosts             int    `json:"num_hosts,omitempty"`
	NumScans             int    `json:"num_scans"`
	NumSessions          int    `json:"num_sessions,omitempty"`
	NumTCPSessions       int    `json:"num_tcp_sessions,omitempty"`
	Owner                string `json:"owner"`
	OwnerID              int    `json:"owner_id"`
	OwnerName            string `json:"owner_name"`
	OwnerUUID            string `json:"owner_uuid"`
	Platform             string `json:"platform,omitempty"`
	Pool                 bool   `json:"pool"`
	ScanCount            int    `json:"scan_count"`
	Shared               int    `json:"shared"`
	Source               string `json:"source"`
	Status               string `json:"status"`
	Timestamp            int    `json:"timestamp"`
	Type                 string `json:"type"`
	UIBuild              string `json:"ui_build,omitempty"`
	UIVersion            string `json:"ui_version,omitempty"`
	UserPermissions      int    `json:"user_permissions"`
	UUID                 string `json:"uuid"`
	AwsUpdateInterval    int    `json:"aws_update_interval,omitempty"`
	License              struct {
		ActivationCode string `json:"activation_code"`
		Users          int    `json:"users"`
		Evaluation     bool   `json:"evaluation"`
		ExpirationDate int    `json:"expiration_date"`
		Agents         int    `json:"agents"`
		Ips            int    `json:"ips"`
		Apps           struct {
			Pci struct {
				Mode string `json:"mode"`
			} `json:"pci"`
			Consec struct {
				Mode           string `json:"mode"`
				ExpirationDate int    `json:"expiration_date"`
			} `json:"consec"`
			Was struct {
				Mode           string `json:"mode"`
				ExpirationDate int    `json:"expiration_date"`
			} `json:"was"`
		} `json:"apps"`
		Scanners     int `json:"scanners"`
		ScannersUsed int `json:"scanners_used"`
		AgentsUsed   int `json:"agents_used"`
	} `json:"license,omitempty"`
}

type AwsTargets struct {
	Targets []struct {
		ContainerUUID string `json:"container_uuid"`
		ScannerUUID   string `json:"scanner_uuid"`
		InstanceID    string `json:"instance_id"`
		PrivateIP     string `json:"private_ip"`
		PublicIP      string `json:"public_ip,omitempty"`
		State         string `json:"state"`
		Zone          string `json:"zone"`
		Type          string `json:"type"`
		Name          string `json:"name,omitempty"`
	} `json:"targets"`
}

type Scanners struct {
	Scanners []Scanner `json:"scanners"`
}

// List scanner instances
func (s *ScannersService) List(ctx context.Context) (*Scanners, *Response, error) {
	scanners := &Scanners{}
	response, err := s.client.Get(ctx, "scanners", nil, scanners)
	return scanners, response, err
}

// List targets for a given AWS scanner
func (s *ScannersService) GetAwsTargets(ctx context.Context, scannerId int) (*AwsTargets, *Response, error) {
	u := fmt.Sprintf("scanners/%d/aws-targets", scannerId)
	targets := &AwsTargets{}
	response, err := s.client.Get(ctx, u, nil, targets)
	return targets, response, err
}
