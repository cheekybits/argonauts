package jsonio

import (
	"bufio"
	"encoding/json"
	"fmt"

	"io"
)

// Writer writes JSON.
type Writer struct {
	e *json.Encoder
	w io.Writer
}

// NewWriter makes a new Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{e: json.NewEncoder(w), w: w}
}

// Write writes the object.
func (w *Writer) Write(o interface{}) error {
	return w.e.Encode(o)
}

// WriteObj writes the object map[string]interface{}.
func (w *Writer) WriteObj(o map[string]interface{}) error {
	return w.Write(o)
}

// WriteArray writes the array []map[string]interface{}.
func (w *Writer) WriteArray(o []map[string]interface{}) error {
	return w.Write(o)
}

// Reader reads JSON.
type Reader struct {
	s    *bufio.Scanner
	last interface{}
	err  error
}

// NewReader makes a new Reader that will read
// from the specified Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		s: bufio.NewScanner(r),
	}
}

// Next reads the next line of JSON.
func (r *Reader) Next() bool {
	if r.Err() != nil {
		return false
	}
	r.last = nil
	return r.s.Scan()
}

// Obj gets the next object or panics if it's not a JSON object.
func (r *Reader) Obj() map[string]interface{} {
	obj, ok := r.ObjOK()
	if !ok {
		panic(r.err.Error())
	}
	return obj
}

// ObjOK gets the next object, or returns false if it cannot.
// Will not panic.
func (r *Reader) ObjOK() (map[string]interface{}, bool) {
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

// Err gets the last error encountered by the Reader.
func (r *Reader) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.s.Err()
}

func panicmsg(obj interface{}) string {
	return fmt.Sprintf("unexpected type %T", obj)
}
