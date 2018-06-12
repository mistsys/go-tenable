package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type WorkbenchesService service

// RESPONSE STRUCTS

type WorkbenchesAssets struct {
	Assets []Asset `json:"assets"`
	Total  int     `json:"total"`
}

type WorkbenchesAssetInfo struct {
	Info struct {
		TimeEnd         time.Time `json:"time_end"`
		TimeStart       time.Time `json:"time_start"`
		ID              string    `json:"id"`
		UUID            string    `json:"uuid"`
		OperatingSystem []string  `json:"operating_system"`
		Fqdn            []string  `json:"fqdn"`
		Counts          struct {
			Vulnerabilities struct {
				Total      int `json:"total"`
				Severities []struct {
					Count int    `json:"count"`
					Level int    `json:"level"`
					Name  string `json:"name"`
				} `json:"severities"`
			} `json:"vulnerabilities"`
			Audits struct {
				Total    int `json:"total"`
				Statuses []struct {
					Count int    `json:"count"`
					Level int    `json:"level"`
					Name  string `json:"name"`
				} `json:"statuses"`
			} `json:"audits"`
		} `json:"counts"`
		HasAgent                  bool      `json:"has_agent"`
		CreatedAt                 time.Time `json:"created_at"`
		UpdatedAt                 time.Time `json:"updated_at"`
		FirstSeen                 time.Time `json:"first_seen"`
		LastSeen                  time.Time `json:"last_seen"`
		LastAuthenticatedScanDate time.Time `json:"last_authenticated_scan_date"`
		LastLicensedScanDate      time.Time `json:"last_licensed_scan_date"`
		Sources                   []struct {
			Name      string    `json:"name"`
			FirstSeen time.Time `json:"first_seen"`
			LastSeen  time.Time `json:"last_seen"`
		} `json:"sources"`
		Tags                    []string `json:"tags"`
		Ipv4                    []string `json:"ipv4"`
		Ipv6                    []string `json:"ipv6"`
		MacAddress              []string `json:"mac_address"`
		NetbiosName             []string `json:"netbios_name"`
		SystemType              []string `json:"system_type"`
		TenableUUID             []string `json:"tenable_uuid"`
		Hostname                []string `json:"hostname"`
		AgentName               []string `json:"agent_name"`
		BiosUUID                []string `json:"bios_uuid"`
		AwsEc2InstanceID        []string `json:"aws_ec2_instance_id"`
		AwsEc2InstanceAmiID     []string `json:"aws_ec2_instance_ami_id"`
		AwsOwnerID              []string `json:"aws_owner_id"`
		AwsAvailabilityZone     []string `json:"aws_availability_zone"`
		AwsRegion               []string `json:"aws_region"`
		AwsVpcID                []string `json:"aws_vpc_id"`
		AwsEc2InstanceGroupName []string `json:"aws_ec2_instance_group_name"`
		AwsEc2InstanceStateName []string `json:"aws_ec2_instance_state_name"`
		AwsEc2InstanceType      []string `json:"aws_ec2_instance_type"`
		AwsSubnetID             []string `json:"aws_subnet_id"`
		AwsEc2ProductCode       []string `json:"aws_ec2_product_code"`
		AwsEc2Name              []string `json:"aws_ec2_name"`
		AzureVMID               []string `json:"azure_vm_id"`
		AzureResourceID         []string `json:"azure_resource_id"`
		SSHFingerprint          []string `json:"ssh_fingerprint"`
		McafeeEpoGUID           []string `json:"mcafee_epo_guid"`
		McafeeEpoAgentGUID      []string `json:"mcafee_epo_agent_guid"`
		QualysAssetID           []string `json:"qualys_asset_id"`
		QualysHostID            []string `json:"qualys_host_id"`
		ServicenowSysid         []string `json:"servicenow_sysid"`
	} `json:"info"`
}

type WorkbenchesAssetVulnerabilities struct {
	Vulnerabilities         []Vulnerability `json:"vulnerabilities"`
	TotalVulnerabilityCount int             `json:"total_vulnerability_count"`
	TotalAssetCount         int             `json:"total_asset_count"`
}

type WorkbenchesAssetVulnerabilityInfo struct {
	Vulnerabilities         []Vulnerability `json:"vulnerabilities"`
	TotalVulnerabilityCount int             `json:"total_vulnerability_count"`
	TotalAssetCount         int             `json:"total_asset_count"`
}

type WorkbenchesVulnerabilities struct {
	Vulnerabilities         []Vulnerability `json:"vulnerabilities"`
	TotalVulnerabilityCount int             `json:"total_vulnerability_count"`
	TotalAssetCount         int             `json:"total_asset_count"`
}

type WorkbenchesVulnerabilityInfo struct {
	// TODO rename
	Info struct {
		Count       int    `json:"count"`
		Description string `json:"description"`
		Synopsis    string `json:"synopsis"`
		Solution    string `json:"solution"`
		Discovery   struct {
			SeenFirst time.Time `json:"seen_first"`
			SeenLast  time.Time `json:"seen_last"`
		} `json:"discovery"`
		Severity      int `json:"severity"`
		PluginDetails struct {
			Family           string    `json:"family"`
			ModificationDate time.Time `json:"modification_date"`
			Name             string    `json:"name"`
			PublicationDate  time.Time `json:"publication_date"`
			Type             string    `json:"type"`
			Version          string    `json:"version"`
			Severity         int       `json:"severity"`
		} `json:"plugin_details"`
		ReferenceInformation []struct {
			Name string `json:"name"`
			URL  string `json:"url,omitempty"`
			// TODO .values is occasionally int, string, int string. This probably happens elsewhere,
			// so that lenient unmarshal in util.go may need to be extended somehow to cover these cases.
			// Or just somehow handle more gracefully
			Values []string `json:"values"`
		} `json:"reference_information"`
		// NOTE api defines these 'interface' fields as just 'object'
		RiskInformation struct {
			RiskFactor          string      `json:"risk_factor"`
			CvssVector          string      `json:"cvss_vector"`
			CvssBaseScore       string      `json:"cvss_base_score"`
			CvssTemporalVector  interface{} `json:"cvss_temporal_vector"`
			CvssTemporalScore   interface{} `json:"cvss_temporal_score"`
			Cvss3Vector         string      `json:"cvss3_vector"`
			Cvss3BaseScore      string      `json:"cvss3_base_score"`
			Cvss3TemporalVector interface{} `json:"cvss3_temporal_vector"`
			Cvss3TemporalScore  interface{} `json:"cvss3_temporal_score"`
			StigSeverity        string      `json:"stig_severity"`
		} `json:"risk_information"`
		SeeAlso []string `json:"see_also"`
		// this name is overloaded
		VulnerabilityInformation struct {
			VulnerabilityPublicationDate time.Time     `json:"vulnerability_publication_date"`
			ExploitedByMalware           interface{}   `json:"exploited_by_malware"`
			PatchPublicationDate         time.Time     `json:"patch_publication_date"`
			ExploitAvailable             interface{}   `json:"exploit_available"`
			ExploitabilityEase           interface{}   `json:"exploitability_ease"`
			AssetInventory               interface{}   `json:"asset_inventory"`
			DefaultAccount               interface{}   `json:"default_account"`
			ExploitedByNessus            interface{}   `json:"exploited_by_nessus"`
			InTheNews                    interface{}   `json:"in_the_news"`
			Malware                      interface{}   `json:"malware"`
			UnsupportedByVendor          interface{}   `json:"unsupported_by_vendor"`
			Cpe                          []string      `json:"cpe"`
			ExploitFrameworks            []interface{} `json:"exploit_frameworks"`
		} `json:"vulnerability_information"`
	} `json:"info"`
}

type WorkbenchesVulnerabilityOutputs struct {
	Outputs []VulnerabilityOutputs `json:"outputs"`
}

func (s *WorkbenchesService) Vulnerabilities(ctx context.Context) (*WorkbenchesVulnerabilities, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, "workbenches/vulnerabilities", nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	props := &WorkbenchesVulnerabilities{}
	response, err := s.client.doRequest(ctx, req, props)
	return props, response, err
}

func (s *WorkbenchesService) VulnerabilityInfo(ctx context.Context, id string) (*WorkbenchesVulnerabilityInfo, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, fmt.Sprintf("workbenches/vulnerabilities/%s/info", id), nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	info := &WorkbenchesVulnerabilityInfo{}
	response, err := s.client.doRequest(ctx, req, info)
	return info, response, err
}

func (s *WorkbenchesService) VulnerabilityOutputs(ctx context.Context, id string) (*WorkbenchesVulnerabilityOutputs, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, fmt.Sprintf("workbenches/vulnerabilities/%s/outputs", id), nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	pluginOutputs := &WorkbenchesVulnerabilityOutputs{}
	response, err := s.client.doRequest(ctx, req, pluginOutputs)
	return pluginOutputs, response, err
}

func (s *WorkbenchesService) Assets(ctx context.Context) (*WorkbenchesAssets, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, "workbenches/assets", nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	assets := &WorkbenchesAssets{}
	response, err := s.client.doRequest(ctx, req, assets)
	return assets, response, err
}

func (s *WorkbenchesService) AssetInfo(ctx context.Context, id string) (*WorkbenchesAssetInfo, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, fmt.Sprintf("workbenches/assets/%s/info", id), nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	info := &WorkbenchesAssetInfo{}
	response, err := s.client.doRequest(ctx, req, info)
	return info, response, err
}

func (s *WorkbenchesService) AssetVulnerabilities(ctx context.Context, id string) (*WorkbenchesAssetVulnerabilities, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, fmt.Sprintf("workbenches/assets/%s/vulnerabilities", id), nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	vulns := &WorkbenchesAssetVulnerabilities{}
	response, err := s.client.doRequest(ctx, req, vulns)
	return vulns, response, err
}

func (s *WorkbenchesService) AssetVulnerabilityInfo(ctx context.Context, assetId string, pluginId string) (*WorkbenchesAssetVulnerabilityInfo, *Response, error) {
	req, err := s.client.createRequest(http.MethodGet, fmt.Sprintf("workbenches/assets/%s/vulnerabilities/%s/info", assetId, pluginId), nil)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "Failed to create request")
	}
	vulns := &WorkbenchesAssetVulnerabilityInfo{}
	response, err := s.client.doRequest(ctx, req, vulns)
	return vulns, response, err
}
