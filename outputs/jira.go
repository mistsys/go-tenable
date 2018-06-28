package outputs

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
)

type CsvAble interface {
	ToCsvHeader() []string
	ToCsvRecords() [][]string
}

type JiraTicket struct {
	Source CsvAble
}

func (j *JiraTicket) Produce() (string, error) {
	var buf bytes.Buffer

	header := j.Source.ToCsvHeader()
	headerLength := len(header)
	records := j.Source.ToCsvRecords()

	writer := csv.NewWriter(&buf)
	if err := writer.Write(header); err != nil {
		return "", errors.New("Failed to write CSV header!")
	}
	for _, record := range records {
		// the csv writer library doesn't check this
		if len(record) != headerLength {
			return "", errors.New(fmt.Sprintf("CSV record and length mismatch: %q", record))
		}
		if err := writer.Write(record); err != nil {
			return "", errors.New("Failed to write CSV!")
		}
		writer.Flush()
	}

	return string(buf.Bytes()), nil
}

/*
const (
	assetVulnerabilitiesTicketTemplate = `not implemented`
	csvTag                             = "csv"
	jiraTicketHeader                   = "Summary,Description,Issue Type,Status,Component"
)

type JiraTicketOpts struct {
	Summary     string `csv:"Summary"`
	Description string `csv:"Description"`
	IssueType   string `csv:"Issue Type"`
	// Status      string `csv:"Status"` // not showing up in the importer for some reason
	Component string `csv:"Component"`
}

// reflect the struct, get all the csv tags, turn them into a csv header by intersposing with ","
// this is not robust in any way
func (ticket *JiraTicketOpts) Header() string {
	var names []string

	t := reflect.Indirect(reflect.ValueOf(ticket)).Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		names = append(names, field.Tag.Get(csvTag))
	}

	return strings.Join(names, ",")
}

func (ticket *JiraTicketOpts) Encode() string {
	var vals []string

	t := reflect.ValueOf(ticket).Elem()
	for i := 0; i < t.NumField(); i++ { // is field order stable?
		field := t.Field(i)
		vals = append(vals, field.String())
	}

	return strings.Join(vals, ",")
}

// XXX I don't know that I like this
func jiraTicket(data interface{}, opts JiraTicketOpts) (string, error) {
	var b bytes.Buffer
	b.WriteString("asdf")

	t, err := template.New("huh").Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = t.Execute(out, data)
	return err
}

JiraTicketOpts defines fixed fields to use in the output CSV; if
unc (v *VulnerabilityInfo) ToJiraTicket(opts *JiraTicketOpts) (string, error) {
 fmt.Println(jiraTicketHeader)
 vulns := v.Vulnerabilities
 for i := 0; i < len(vulns); i++ {
	vuln := vulns[i]
	fmt.Println(vuln.PluginFamily)
 }
 // return toJiraTicket(a, assetVulnerabilitiesTicketTemplate)
 return "", nil
*/
