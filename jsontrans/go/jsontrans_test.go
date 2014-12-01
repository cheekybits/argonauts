package jsontrans_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cheekybits/argonauts/jsontrans/go"
	"github.com/cheekybits/is"
)

func TestTransformer(t *testing.T) {
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
	} {

		in := strings.NewReader(test.S)
		var actualbuf bytes.Buffer
		actual := &actualbuf
		j := jsontrans.New(in, actual, test.A)
		is.NoErr(j.Go())
		is.Equal(strings.TrimRight(actual.String(), "\n"), test.E)

	}

}
