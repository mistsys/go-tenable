package client

import (
	"time"
)

type Asset struct {
	ID       string    `json:"id"`
	HasAgent bool      `json:"has_agent"`
	LastSeen time.Time `json:"last_seen"`
	Sources  []struct {
		Name      string    `json:"name"`
		FirstSeen time.Time `json:"first_seen"`
		LastSeen  time.Time `json:"last_seen"`
	} `json:"sources"`
	// NOTE these types are just observed... the API docs don't specify a type
	Ipv4            []string `json:"ipv4"`
	Ipv6            []string `json:"ipv6"`
	Fqdn            []string `json:"fqdn"`
	NetbiosName     []string `json:"netbios_name"`
	OperatingSystem []string `json:"operating_system"`
	AgentName       []string `json:"agent_name"`
	MacAddress      []string `json:"mac_address"`
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
	Count              int    `json:"count"`
	PluginFamily       string `json:"plugin_family"`
	PluginID           int    `json:"plugin_id"`
	PluginName         string `json:"plugin_name"`
	VulnerabilityState string `json:"vulnerability_state"`
	AcceptedCount      int    `json:"accepted_count"`
	RecastedCount      int    `json:"recasted_count"`
	CountsBySeverity   []struct {
		Count int `json:"count"`
		Value int `json:"value"`
	} `json:"counts_by_severity"`
	Severity int `json:"severity"`
}

// what's this
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
	Vulns       int    `json:"vulns"`
}
