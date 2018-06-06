package client

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

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

func (t *TenableClient) ScansList(ctx context.Context) (*ScansList, error) {
	req, err := t.createRequest(http.MethodGet, "scans", nil)
	if err != nil {
		log.Printf("Failed to create request %s", err)
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	list := &ScansList{}
	err = t.doRequest(ctx, req, list)
	return list, err
}

func (t *TenableClient) ScanDetail(ctx context.Context, id string) (*ScanDetail, error) {
	req, err := t.createRequest(http.MethodGet, fmt.Sprintf("scans/%s", id), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create request")
	}
	status := &ScanDetail{}
	err = t.doRequest(ctx, req, status)
	return status, err
}
