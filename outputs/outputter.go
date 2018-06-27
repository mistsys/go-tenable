package outputs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func jiraProcessor(data interface{}) (string, error) {
	if ticketSource, ok := data.(CsvAble); ok == true {
		ticket := &JiraTicket{Source: ticketSource}
		csvString, err := ticket.Produce()
		// not sure if should return error at all
		if err != nil {
			fmt.Printf("%v", err)
			return "", errors.Wrap(err, "Error producing CSV")
		}
		return csvString, err
	}
	return "", errors.New("This datatype does not support JIRA CSV output")
}

// FIXME
// could I make default in the cmd actually be something like *raw* and pass the response.BodyJson() through?
// then Output takes some sort of raw as an input, or it gets tagged onto the input struct somehow
// point is I want outputter to take *one* input in a way that requires no logic on the calling side
func jsonProcessor(data interface{}) (string, error) {
	// XXX if the data is a string, just shoot it back. this is... not good
	switch v := data.(type) {
	case string:
		return v, nil
	}

	s, err := json.MarshalIndent(data, "  ", "  ") // XXX repeated work!
	if err != nil {
		return "", errors.Wrap(err, "Failed to JSON format data")
	}
	// can this just return bytes?
	return string(s), err
}

type Outputter struct {
	Verbose     bool
	Format      string
	Destination io.Writer
}

func NewOutputter(verbose bool, format string, destination io.Writer) *Outputter {
	return &Outputter{
		Verbose:     verbose,
		Format:      format,
		Destination: destination,
	}
}

// return options are string, error; error; string?
// if this function is meant to consolidate all outputting functionality, then this thing should definitely handle errors itself
func (o *Outputter) Output(data interface{}) error {
	var msg string
	var err error

	// other ways?
	switch o.Format {
	// I don't like that the Output function defines processors this way; I'd *prefer* that the "constructor" takes in a map
	// of processors that can be accessed at runtime
	case "jira":
		msg, err = jiraProcessor(data)
	case "json":
		msg, err = jsonProcessor(data)
	case "raw": // see above?
		msg, err = jsonProcessor(data)
	// case "template":
	// 	msg, err = templateProcessor(data)
	default:
		msg, err = jsonProcessor(data)
	}

	if err != nil {
		fmt.Printf("Error producing output: %v", err)
		return err
	} else {
		// FIXME
		_, err := o.Destination.Write([]byte(msg))
		return err // XXX
	}
}

// ugly!
func (o *Outputter) SetFormat(format string) {
	o.Format = format
}

// cowardly newfile function which creates a new file or errors if it already exists
func NewFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		// file exists!
		return nil, errors.New(fmt.Sprintf("File %s already exists", filename))
	}

	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return f, err
}
