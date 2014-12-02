package jsontrans_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cheekybits/argonauts/jsontrans/go"
	"github.com/cheekybits/is"
)

func TestLower(t *testing.T) {
	is := is.New(t)

	for _, test := range []struct {
		A []string
		S string
		E string
	}{
		{
			[]string{"-lower"},
			`{"Field":"Value"}`,
			`{"field":"Value"}`,
		},
		{
			[]string{"-array=2"},
			`{"item":1}
{"item":2}
{"item":3}
{"item":4}
{"item":5}
{"item":6}
{"item":7}
{"item":8}
{"item":9}
{"item":10}
`,
			`[{"item":1},{"item":2}]
[{"item":3},{"item":4}]
[{"item":5},{"item":6}]
[{"item":7},{"item":8}]
[{"item":9},{"item":10}]`,
		},
	} {

		in := strings.NewReader(test.S)
		var actualbuf bytes.Buffer
		actual := &actualbuf
		j := jsontrans.New(in, actual, test.A)
		is.NoErr(j.Go())
		is.Equal(strings.TrimRight(actual.String(), "\n"), test.E)

	}

}
