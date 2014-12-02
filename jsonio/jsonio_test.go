package jsonio_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cheekybits/argonauts/jsonio"
	"github.com/cheekybits/is"
)

func TestReader(t *testing.T) {
	is := is.New(t)

	in := strings.NewReader(`{"name":"cheekybits"}
{"name":"matryer"}
{"name":"tylerb"}`)
	r := jsonio.NewReader(in)

	var objs []map[string]interface{}
	for r.Next() {
		o, ok := r.ObjOK()
		is.Equal(ok, true)
		is.Equal(r.Obj(), o)
		objs = append(objs, o)
	}
	is.NoErr(r.Err())
	is.Equal(len(objs), 3)

	is.Equal(objs[0]["name"], "cheekybits")
	is.Equal(objs[1]["name"], "matryer")
	is.Equal(objs[2]["name"], "tylerb")

}

func TestReaderErr(t *testing.T) {
	is := is.New(t)

	in := strings.NewReader(`{"name":"cheekybits"}
{"name":"matryer"}
{"name":"tylerb"`)
	r := jsonio.NewReader(in)

	var oks []bool
	var objs []map[string]interface{}
	for r.Next() {
		o, ok := r.ObjOK()
		oks = append(oks, ok)
		objs = append(objs, o)
	}
	is.OK(r.Err())
	is.Equal(r.Err().Error(), "unexpected end of JSON input")

}

func TestWriter(t *testing.T) {
	is := is.New(t)

	var buf bytes.Buffer
	jsonbuf := &buf

	w := jsonio.NewWriter(jsonbuf)
	w.Write("Mat")
	is.Equal(jsonbuf.String(), `"Mat"`+"\n")

	buf.Reset()
	w.WriteObj(map[string]interface{}{"name": "Mat"})
	is.Equal(jsonbuf.String(), `{"name":"Mat"}`+"\n")

	buf.Reset()
	w.WriteArray([]map[string]interface{}{{"name": "Mat"}, {"name": "Tyler"}})
	is.Equal(jsonbuf.String(), `[{"name":"Mat"},{"name":"Tyler"}]`+"\n")

}
