package tenable

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type ServerStatus struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type Folder struct {
	Custom      int    `json:"custom"`
	DefaultTag  int    `json:"default_tag"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UnreadCount int    `json:"unread_count"`
}

// NumericBool type because Tenable sometimes returns
// 1 for what should be boolean
type NumericBool bool

func (n NumericBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n NumericBool) UnmarshalJSON(b []byte) error {
	strB := string(b)
	switch strB {
	case "true":
		n = true
	case "false":
		n = false
	default:
		i, err := strconv.Atoi(string(b))
		if err != nil {
			return errors.Wrapf(err, "Failed to parse as int")
		}
		if i > 0 {
			n = true
		} else {
			n = false
		}
	}
	return nil
}

// End NumericBool

type Scan struct {
	Control              bool        `json:"control"`
	CreationDate         int         `json:"creation_date"`
	Enabled              bool        `json:"enabled"`
	ID                   int         `json:"id"`
	LastModificationDate int         `json:"last_modification_date"`
	Legacy               bool        `json:"legacy"`
	Name                 string      `json:"name"`
	Owner                string      `json:"owner"`
	Permissions          int         `json:"permissions"`
	Read                 NumericBool `json:"read"`
	Rrules               string      `json:"rrules"`
	ScheduleUUID         string      `json:"schedule_uuid"`
	Shared               bool        `json:"shared"`
	Starttime            string      `json:"starttime"`
	Status               string      `json:"status"`
	Timezone             string      `json:"timezone"`
	Type                 string      `json:"type"`
	UserPermissions      int         `json:"user_permissions"`
	UUID                 string      `json:"uuid"`
}

type ScansList struct {
	Folders   []Folder `json:"folders"`
	Scans     []Scan   `json:"scans"`
	Timestamp int      `json:"timestamp"`
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

type Host struct {
	AssetID             int    `json:"asset_id"`
	Critical            int    `json:"critical"`
	High                int    `json:"high"`
	HostID              int    `json:"host_id"`
	HostIndex           int    `json:"host_index"`
	Hostname            string `json:"hostname"`
	Info                int    `json:"info"`
	Low                 int    `json:"low"`
	Medium              int    `json:"medium"`
	Numchecksconsidered int    `json:"numchecksconsidered"`
	Progress            string `json:"progress"`
	Scanprogresscurrent int    `json:"scanprogresscurrent"`
	Scanprogresstotal   int    `json:"scanprogresstotal"`
	Score               int    `json:"score"`
	Severity            int    `json:"severity"`
	Severitycount       struct {
		Item []struct {
			Count         int `json:"count"`
			Severitylevel int `json:"severitylevel"`
		} `json:"item"`
	} `json:"severitycount"`
	Totalchecksconsidered int `json:"totalchecksconsidered"`
}

type History struct {
	AltTargetsUsed       bool   `json:"alt_targets_used"`
	CreationDate         int    `json:"creation_date"`
	HistoryID            int    `json:"history_id"`
	LastModificationDate int    `json:"last_modification_date"`
	OwnerID              int    `json:"owner_id"`
	Scheduler            int    `json:"scheduler"`
	Status               string `json:"status"`
	Type                 string `json:"type"`
	UUID                 string `json:"uuid"`
}

type Vulnerability struct {
	Count        int    `json:"count"`
	PluginFamily string `json:"plugin_family"`
	PluginID     int    `json:"plugin_id"`
	PluginName   string `json:"plugin_name"`
	Severity     int    `json:"severity"`
	VulnIndex    int    `json:"vuln_index"`
}

type Filter struct {
	Control struct {
		ReadableRegex string `json:"readable_regex"`
		Regex         string `json:"regex"`
		Type          string `json:"type"`
	} `json:"control"`
	GroupName    string   `json:"group_name"`
	Name         string   `json:"name"`
	Operators    []string `json:"operators"`
	ReadableName string   `json:"readable_name"`
}

type Note struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Severity int    `json:"severity"`
}

type Remediation struct {
	Value       string `json:"value"`
	Remediation string `json:"remediation"`
	Hosts       int    `json:"hosts"`
	vulns       int    `json:"vulns"`
}

type ScanDetail struct {
	Comphosts  []Host          `json:"comphosts"`
	Compliance []Vulnerability `json:"compliance"`
	Filters    []Filter        `json:"filters"`
	History    []History       `json:"history"`
	Hosts      []Host          `json:"hosts"`
	Info       struct {
		Acls []struct {
			DisplayName interface{} `json:"display_name"`
			ID          interface{} `json:"id"`
			Name        interface{} `json:"name"`
			Owner       interface{} `json:"owner"`
			Permissions int         `json:"permissions"`
			Type        string      `json:"type"`
		} `json:"acls"`
		AltTargetsUsed  bool        `json:"alt_targets_used"`
		Control         bool        `json:"control"`
		EditAllowed     bool        `json:"edit_allowed"`
		FolderID        int         `json:"folder_id"`
		Hasaudittrail   bool        `json:"hasaudittrail"`
		Haskb           bool        `json:"haskb"`
		Hostcount       int         `json:"hostcount"`
		Name            string      `json:"name"`
		NoTarget        bool        `json:"no_target"`
		ObjectID        int         `json:"object_id"`
		Owner           string      `json:"owner"`
		Pci_can_upload  bool        `json:"pci-can-upload"`
		Policy          string      `json:"policy"`
		ScanEnd         int         `json:"scan_end"`
		ScanStart       int         `json:"scan_start"`
		ScanType        string      `json:"scan_type"`
		ScannerEnd      interface{} `json:"scanner_end"`
		ScannerName     string      `json:"scanner_name"`
		ScannerStart    interface{} `json:"scanner_start"`
		ScheduleUUID    string      `json:"schedule_uuid"`
		Shared          interface{} `json:"shared"`
		Status          string      `json:"status"`
		Targets         string      `json:"targets"`
		Timestamp       int         `json:"timestamp"`
		UserPermissions int         `json:"user_permissions"`
		UUID            string      `json:"uuid"`
	} `json:"info"`
	Notes        []Note `json:"notes"`
	Remediations struct {
		NumCves           int           `json:"num_cves"`
		NumHosts          int           `json:"num_hosts"`
		NumImpactedHosts  int           `json:"num_impacted_hosts"`
		NumRemediatedCves int           `json:"num_remediated_cves"`
		Remediations      []Remediation `json:"remediations"`
	} `json:"remediations"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}
