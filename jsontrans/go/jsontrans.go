package jsontrans

import (
	"flag"
	"io"
	"strings"

	"github.com/cheekybits/argonauts/jsonio"
)

// Transformer can transform JSON.
type Transformer struct {
	r           *jsonio.Reader
	w           *jsonio.Writer
	LowerFields bool
	ArraySize   int
}

// New makes a new Transformer.
func New(in io.Reader, out io.Writer, args []string) *Transformer {
	fs := flag.NewFlagSet("name", flag.ExitOnError)
	var (
		lowerfields = fs.Bool("lower", false, "make fields lowercase")
		arraySize   = fs.Int("array", 0, "group objects into arrays of this size")
	)
	fs.Parse(args)
	return &Transformer{
		r:           jsonio.NewReader(in),
		w:           jsonio.NewWriter(out),
		LowerFields: *lowerfields,
		ArraySize:   *arraySize,
	}
}

// Go processes the input and writes the transofrmed JSON
// to the output.
func (j *Transformer) Go() error {
	if j.ArraySize > 0 {

		arr := make([]map[string]interface{}, j.ArraySize)
		i := 0
		for j.r.Next() {
			obj, ok := j.r.ObjOK()
			if !ok {
				return j.r.Err()
			}
			obj = j.Transform(obj)
			arr[i] = obj
			if i == j.ArraySize-1 {
				// full buffer
				if err := j.w.WriteArray(arr); err != nil {
					return err
				}
				i = -1
			}
			i++
		}

		return nil
	}
	for j.r.Next() {
		obj, ok := j.r.ObjOK()
		if !ok {
			return j.r.Err()
		}
		obj = j.Transform(obj)
		if err := j.w.WriteObj(obj); err != nil {
			return err
		}
	}
	return j.r.Err()
}

// Transform transforms an individual object.
func (j *Transformer) Transform(m map[string]interface{}) map[string]interface{} {
	n := make(map[string]interface{})
	for k, v := range m {
		if j.LowerFields {
			k = strings.ToLower(string(k[0])) + string(k[1:])
		}
		n[k] = v
	}
	return n
}
