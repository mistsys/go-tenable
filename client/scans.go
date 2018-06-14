package client

import (
	"context"
	"fmt"
	"net/http"
)

type ScansService service

// RESPONSE STRUCTS

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

type ScansList struct {
	Folders   []Folder `json:"folders"`
	Scans     []Scan   `json:"scans"`
	Timestamp int      `json:"timestamp"`
}

// REQUEST STRUCTS

// The full flow for downloading a scan report is:
// 1. hit /scans/export-request to start preparing a report file
// 2. poll /scans/export-status until the file is ready
// 3. hit /scans/export-download to get the report file

// response object when you successfully *request* a download
type ScanExportRequest struct {
	File      string `json:"file"`
	TempToken string `json:"temp_token"`
}

type ScanExportStatus struct {
	foo string
}

type ScanExportDownload struct {
	foo string
}

func (s *ScansService) List(ctx context.Context) (*ScansList, *Response, error) {
	list := &ScansList{}
	response, err := s.client.Get(ctx, "scans", nil, nil, list)
	return list, response, err
}

func (s *ScansService) Detail(ctx context.Context, scanId string) (*ScanDetail, *Response, error) {
	u := fmt.Sprintf("scans/%s", scanId)
	status := &ScanDetail{}
	response, err := s.client.Get(ctx, u, nil, nil, status)
	return status, response, err
}

// TODO actually supposed to be a POST
func (s *ScansService) ExportRequest(ctx context.Context, scanId string) (*ScanDetail, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("scans/%s", scanId), nil)
	if err != nil {
		return nil, nil, err
	}
	status := &ScanDetail{}
	response, err := s.client.Do(ctx, req, status)
	return status, response, err
}
