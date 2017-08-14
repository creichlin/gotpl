package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Reads a YAML document from the values_in stream, uses it as values
// for the tpl_files templates and writes the executed templates to
// the out stream.
func ExecuteTemplates(values_in io.Reader, out io.Writer, t string) error {
	tpl := template.New("test")
	tpl, err := tpl.Funcs(template.FuncMap{
		"prefixLines": prefixLines,
		"suffixLines": suffixLines,
	}).Parse(t)
	if err != nil {
		return fmt.Errorf("Error parsing template(s): %v", err)
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, values_in)
	if err != nil {
		return fmt.Errorf("Failed to read standard input: %v", err)
	}

	var values map[string]interface{}
	err = yaml.Unmarshal(buf.Bytes(), &values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}

	err = tpl.Execute(out, values)
	if err != nil {
		return fmt.Errorf("Failed to parse standard input: %v", err)
	}
	return nil
}

func main() {
	ct, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Printf("Failed to read file %v", os.Args[1:])
		os.Exit(1)
	}
	err = ExecuteTemplates(os.Stdin, os.Stdout, string(ct))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
