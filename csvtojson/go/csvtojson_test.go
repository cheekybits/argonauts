package csvtojson_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cheekybits/argonauts/csvtojson/go"
	"github.com/cheekybits/is"
)

func TestGo(t *testing.T) {
	is := is.New(t)

	in := strings.NewReader(`"Name","Number","Another"
"Mat",100,"thing"
"Tyler",101,"thing"
"Ryan",102,""
"",103,"thing"
"Another",,"thing"`)
	var outbuf bytes.Buffer
	out := &outbuf

	c := csvtojson.New(in, out)
	is.NoErr(c.Go())

	var items []map[string]interface{}
	var err error
	s := bufio.NewScanner(out)
	for s.Scan() {
		data := make(map[string]interface{})
		if err = json.NewDecoder(bytes.NewReader(s.Bytes())).Decode(&data); err != nil {
			break
		}
		items = append(items, data)
	}
	is.NoErr(err)

	is.Equal(len(items), 5)
	is.Equal(items[0]["Name"], "Mat")
	is.Equal(items[0]["Number"], 100)
	is.Equal(items[0]["Another"], "thing")

}
