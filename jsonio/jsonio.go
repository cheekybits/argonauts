package jsonio

import (
	"bufio"
	"encoding/json"
	"fmt"

	"io"
)

// JSONWriter writes JSON.
type JSONWriter struct {
	e *json.Encoder
}

// NewWriter makes a new JSONWriter.
func NewWriter(w io.Writer) *JSONWriter {
	return &JSONWriter{e: json.NewEncoder(w)}
}

// Write writes the object.
func (w *JSONWriter) Write(o interface{}) error {
	return w.e.Encode(o)
}

// WriteObj writes the map[string]interface{}.
func (w *JSONWriter) WriteObj(o map[string]interface{}) error {
	return w.Write(o)
}

// JSONReader reads JSON.
type JSONReader struct {
	s    *bufio.Scanner
	last interface{}
	err  error
}

// NewReader makes a new JSONReader that will read
// from the specified Reader.
func NewReader(r io.Reader) *JSONReader {
	return &JSONReader{
		s: bufio.NewScanner(r),
	}
}

// Next reads the next line of JSON.
func (r *JSONReader) Next() bool {
	if r.Err() != nil {
		return false
	}
	r.last = nil
	return r.s.Scan()
}

// Obj gets the next object or panics if it's not a JSON object.
func (r *JSONReader) Obj() map[string]interface{} {
	obj, ok := r.ObjOK()
	if !ok {
		panic(r.err.Error())
	}
	return obj
}

// ObjOK gets the next object, or returns false if it cannot.
// Will not panic.
func (r *JSONReader) ObjOK() (map[string]interface{}, bool) {
	if r.last == nil {
		if r.err = json.Unmarshal(r.s.Bytes(), &r.last); r.err != nil {
			return nil, false
		}
	}
	lastmap, ok := r.last.(map[string]interface{})
	if !ok {
		r.err = fmt.Errorf("unexpected type %T", r.last)
	}
	return lastmap, ok
}

// Err gets the last error encountered by the JSONReader.
func (r *JSONReader) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.s.Err()
}

func panicmsg(obj interface{}) string {
	return fmt.Sprintf("unexpected type %T", obj)
}
