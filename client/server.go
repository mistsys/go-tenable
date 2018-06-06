package client

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type ServerService service

type ServerStatus struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ServerProperties struct {
	Analytics struct {
		Enabled bool   `json:"enabled"`
		Key     string `json:"key"`
		SiteID  string `json:"site_id"`
	} `json:"analytics"`
	Capabilities struct {
		MultiScanner      bool   `json:"multi_scanner"`
		MultiUser         string `json:"multi_user"`
		ReportEmailConfig bool   `json:"report_email_config"`
		TwoFactor         struct {
			SMTP   bool `json:"smtp"`
			Twilio bool `json:"twilio"`
		} `json:"two_factor"`
	} `json:"capabilities"`
	ContainerDbVersion string `json:"container_db_version"`
	Enterprise         bool   `json:"enterprise"`
	Evaluation         struct {
		LimitEnabled bool `json:"limitEnabled"`
		Scans        int  `json:"scans"`
		Targets      int  `json:"targets"`
	} `json:"evaluation"`
	Expiration     int    `json:"expiration"`
	ExpirationTime int    `json:"expiration_time"`
	ForceUIReload  bool   `json:"force_ui_reload"`
	IdleTimeout    string `json:"idle_timeout"`
	License        struct {
		ActivationCode string `json:"activation_code"`
		Agents         int    `json:"agents"`
		AgentsUsed     int    `json:"agents_used"`
		Apps           struct {
			Consec struct {
				ExpirationDate int    `json:"expiration_date"`
				Mode           string `json:"mode"`
			} `json:"consec"`
			Pci struct {
				Mode string `json:"mode"`
			} `json:"pci"`
			Was struct {
				ExpirationDate int    `json:"expiration_date"`
				Mode           string `json:"mode"`
			} `json:"was"`
		} `json:"apps"`
		Evaluation     bool `json:"evaluation"`
		ExpirationDate int  `json:"expiration_date"`
		Ips            int  `json:"ips"`
		Scanners       int  `json:"scanners"`
		ScannersUsed   int  `json:"scanners_used"`
		Users          int  `json:"users"`
	} `json:"license"`
	LimitEnabled    bool          `json:"limitEnabled"`
	LoadedPluginSet string        `json:"loaded_plugin_set"`
	LoginBanner     interface{}   `json:"login_banner"`
	Msp             bool          `json:"msp"`
	NessusType      string        `json:"nessus_type"`
	NessusUIBuild   string        `json:"nessus_ui_build"`
	NessusUIVersion string        `json:"nessus_ui_version"`
	Notifications   []interface{} `json:"notifications"`
	PluginSet       string        `json:"plugin_set"`
	ScannerBoottime int           `json:"scanner_boottime"`
	ServerBuild     string        `json:"server_build"`
	ServerUUID      string        `json:"server_uuid"`
	ServerVersion   string        `json:"server_version"`
	Update          struct {
		Href       interface{} `json:"href"`
		NewVersion int         `json:"new_version"`
		Restart    int         `json:"restart"`
	} `json:"update"`
}

func (s *ServerService) Properties(ctx context.Context) (*ServerProperties, error) {
	req, err := s.client.createRequest(http.MethodGet, "server/properties", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	props := &ServerProperties{}
	err = s.client.doRequest(ctx, req, props)
	return props, err
}

func (s *ServerService) Status(ctx context.Context) (*ServerStatus, error) {
	req, err := s.client.createRequest(http.MethodGet, "server/status", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &ServerStatus{}
	err = s.client.doRequest(ctx, req, status)
	return status, err
}
