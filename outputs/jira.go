package outputs

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
)

type CsvMapReader struct {
    Reader *csv.Reader
    Columns map[string]int
}

type Record struct {
    Reader *CsvMapReader
    Values []string
}

func NewCsvMapReader(r io.Reader) *CsvMapReader {
    return &CsvMapReader{
        Reader: csv.NewReader(r),
        Columns: make(map[string]int),
    }
}

// set the value of c.Columns to look up later
func (c *CsvMapReader) InitColumns() error {
    // should be the csv first line, though there're obviously no guards to prevent
    // you from advancing the reader before this call
    columns, err := c.Reader.Read()
    if err != nil {
        return err
    }
    for n, name := range(columns) {
        c.Columns[name] = n
    }
    return nil
}

func (c *CsvMapReader) Read() (*Record, error) {
    line, err := c.Reader.Read()
    if err != nil {
        return nil, err
    }
    return &Record{
        Reader: c, // so we can refer to columns map later
        Values: line,
    }, nil
}

// Return the column value by name
func (r *Record) GetColumn(name string) string {
    index := r.Reader.Columns[name]
    return r.Values[index] // handle index out of bounds
}

func (r *Record) GetColumns(columns []string) []string {
    var ret []string
    for _, column := range(columns) {
        value := r.GetColumn(column)
        ret = append(ret, value)
    }
    return ret
}

func (r *Record) ToJira() []string {
    host := r.GetColumn("Host")
    synopsis := r.GetColumn("Synopsis")
    description := r.GetColumn("Description")
    issueType := "Bug"
    status := "Open"
    summary := fmt.Sprintf("%s exposes a known vulnerability: %s", host, synopsis)
    ticket := []string{summary, description, issueType, status}

    return ticket
}

// take tenable csv export -> jira tickets
// tenable produces scan export CSVs with these columns (export a file for yourself to see):
// Plugin ID,CVE,CVSS,Risk,Host,Protocol,Port,Name,Synopsis,Description,Solution,See Also,Plugin Output,Asset UUID,Vulnerability State,IP Address,FQDN,NetBios,OS,MAC Address,Plugin Family,CVSS Base Score,CVSS Temporal Score,CVSS Temporal Vector,CVSS Vector,CVSS3 Base Score,CVSS3 Temporal Score,CVSS3 Temporal Vector,CVSS3 Vector,System Type,Host Start,Host End
// future: shouldn't build the whole csv in memory, better to return something like an io.Reader that produces on the fly. That's a thing you can do with go interfaces, right?
var defaultJiraHeader []string = []string{"Summary", "Description", "Issue Type", "Status"}
func WriteTenableToJira(in io.Reader, out io.Writer) error {
    reader := NewCsvMapReader(in)
    err := reader.InitColumns()
    if err != nil {
        return err
    }

    writer := csv.NewWriter(out)
	
	if err := writer.Write(defaultJiraHeader); err != nil {
        return errors.New("Failed to write CSV header!")
    }

	writer.Flush()
    for {
        record, err := reader.Read()
        if err != nil {
            if err == io.EOF {
                break
            }
            return err
        }
        risk := record.GetColumn("Risk")
        if risk != "None" {
            ticket := record.ToJira()
            if err := writer.Write(ticket); err != nil {
                return errors.New("Failed to write CSV!")
            }
            writer.Flush()
        }
    }

    return nil
}
