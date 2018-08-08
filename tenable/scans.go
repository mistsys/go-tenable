package tenable

import (
    "context"
    "fmt"
    "strings"
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

type Scans struct {
    Folders   []Folder `json:"folders"`
    Scans     []Scan   `json:"scans"`
    Timestamp int      `json:"timestamp"`
}

// response when a scan is launched
type ScansLaunch struct {
    ScanUUID string `json:"scan_uuid"`
}

type ScanExportOptions struct {
    ScanId int
    Format string
}

type ScansExportRequest struct {
    File      int    `json:"file"`
    TempToken string `json:"temp_token"`
}

type ScansExportStatus struct {
    Status string `json:"status"`
}

type ScansCreate struct {
    Settings struct {
        Name        string `json:"name"`
        Description string `json:"description"`
        PolicyId    string `json:"policy_id"`
        FolderId    string `json:"folder_id"`
        ScannerId   string `json:"scanner_id"`
        Enabled     string `json:"enabled"`
        Launch      string `json:"launch"`
        Starttime   string `json:"starttime"`
        RRules      string `json:"rrules"`
        Timezone    string `json:"timezone"`
        TextTargets string `json:"text_targets"`
        FileTargets string `json:"file_targets"`
        Emails      string `json:"emails"`
        ACLs        string `json:"acls"`
    } `json:"settings"`
}

func (s *ScansService) List(ctx context.Context) (*Scans, *Response, error) {
    list := &Scans{}
    response, err := s.client.Get(ctx, "scans", nil, list)
    return list, response, err
}

func (s *ScansService) Detail(ctx context.Context, scanId int) (*ScanDetail, *Response, error) {
    u := fmt.Sprintf("scans/%d", scanId)
    status := &ScanDetail{}
    response, err := s.client.Get(ctx, u, nil, status)
    return status, response, err
}

func (s *ScansService) Launch(ctx context.Context, scanId int, targets []string) (*ScansLaunch, *Response, error) {
    u := fmt.Sprintf("scans/%d/launch", scanId)
    launch := &ScansLaunch{}
    response, err := s.client.Post(ctx, u, nil, nil, launch)
    return launch, response, err
}

// TODO pause/resume/stop don't have explicit return specs, just http codes (which, honestly, i prefer)
func (s *ScansService) Pause(ctx context.Context, scanId int, targets []string) (*ScansLaunch, *Response, error) {
    u := fmt.Sprintf("scans/%d/pause", scanId)
    launch := &ScansLaunch{}
    response, err := s.client.Post(ctx, u, nil, nil, launch)
    return launch, response, err
}

func (s *ScansService) Resume(ctx context.Context, scanId int, targets []string) (*ScansLaunch, *Response, error) {
    u := fmt.Sprintf("scans/%d/resume", scanId)
    launch := &ScansLaunch{}
    response, err := s.client.Post(ctx, u, nil, nil, launch)
    return launch, response, err
}

func (s *ScansService) Stop(ctx context.Context, scanId int, targets []string) (*ScansLaunch, *Response, error) {
    u := fmt.Sprintf("scans/%d/stop", scanId)
    launch := &ScansLaunch{}
    response, err := s.client.Post(ctx, u, nil, nil, launch)
    return launch, response, err
}

func (s *ScansService) ExportRequest(ctx context.Context, scanId int, format string) (*ScansExportRequest, *Response, error) {
    u := fmt.Sprintf("scans/%d/export", scanId)
    body := fmt.Sprintf(`{"format":"%s"}`, format) // YIKES
    exportRequest := &ScansExportRequest{}
    response, err := s.client.Post(ctx, u, nil, strings.NewReader(body), exportRequest)
    return exportRequest, response, err
}

// I don't know why the types or responses of the export endpoints are different between workbenches and scans. Ask Tenable
func (s *ScansService) ExportStatus(ctx context.Context, scanId int, fileId int) (*ScansExportStatus, *Response, error) {
    u := fmt.Sprintf("scans/%d/export/%d/status", scanId, fileId)
    exportStatus := &ScansExportStatus{}
    response, err := s.client.Get(ctx, u, nil, exportStatus)
    return exportStatus, response, err
}
