package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlTemplate(t *testing.T) {
	type io struct {
		Input    string
		Template string
		Output   string
	}

	tests := []io{
		{
			Input:    "test: value",
			Template: "{{.test}}",
			Output:   "value",
		},
		{
			Input:    "name: Max\nage: 15",
			Template: "Hello {{.name}}, of {{.age}} years old",
			Output:   "Hello Max, of 15 years old",
		},
		{
			Input:    "legumes:\n  - potato\n  - onion\n  - cabbage",
			Template: "Legumes:{{ range $index, $el := .legumes}}{{if $index}},{{end}} {{$el}}{{end}}",
			Output:   "Legumes: potato, onion, cabbage",
		},
		{
			Input: `itchy: |
  AAAA
  BBBB
  CCCC`,
			Template: "{{.itchy}}",
			Output: `AAAA
BBBB
CCCC`,
		},
		{
			Input:    `itchy: "AAAA\nBBBB\nCCCC"`,
			Template: "{{ .itchy | prefixLines \"  \" }}",
			Output: `  AAAA
  BBBB
  CCCC`,
		},
	}

	for _, test := range tests {
		output := bytes.NewBuffer(nil)
		err := ExecuteTemplates(strings.NewReader(test.Input), output,
			test.Template)
		assert.Nil(t, err)

		assert.Equal(t, test.Output, output.String())

	}
}
