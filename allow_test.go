package http405

import (
	"testing"
)

func TestAllow(t *testing.T) {

	tests := []struct{
		Methods []string
		Expected string
	}{
		{
			Methods: []string{},
		},




		{
			Methods: []string{"GET"},
			Expected: "GET",
		},
		{
			Methods: []string{"HEAD"},
			Expected: "HEAD",
		},
		{
			Methods: []string{"POST"},
			Expected: "POST",
		},



		{
			Methods: []string{"KICK","PUNCH"},
			Expected: "KICK, PUNCH",
		},



		{
			Methods: []string{"GET","HEAD","POST"},
			Expected: "GET, HEAD, POST",
		},



		{
			Methods: []string{"GET,HEAD,POST"},
			Expected: "GET, HEAD, POST",
		},
		{
			Methods: []string{"POST,HEAD,GET"},
			Expected: "GET, HEAD, POST",
		},
	}

	for testNumber, test := range tests {

		actual := allow(test.Methods...)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'allow' HTTP-header value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}

	}
}
