package csvtojson

import (
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/cheekybits/parse"
)

// Converter converts from CSV inupt, to lines of JSON output.
type Converter struct {
	in  io.Reader
	out io.Writer
}

// New creates a new Converter that will read CSV input from in,
// and write JSON output to out.
func New(in io.Reader, out io.Writer) *Converter {
	return &Converter{in: in, out: out}
}

// Go processes the input and generates the output.
func (c *Converter) Go() error {

	r := csv.NewReader(c.in)
	head, err := r.Read()

	j := json.NewEncoder(c.out)

	var row []string
	for {
		row, err = r.Read()
		if err != nil {
			break
		}
		data := make(map[string]interface{})
		for i, h := range head {
			data[h] = parse.String(row[i])
		}
		if err = j.Encode(data); err != nil {
			break
		}
	}

	if err != io.EOF && err != nil {
		return err
	}

	return nil
}
