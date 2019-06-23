package report

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func Generate(filename string) error {
	file, err := readFile(filename)
	if err != nil {
		return err
	}

	root := parseMarkdown(file)
	if _, ok := root.(*ast.Document); ok {
		fmt.Println("found the root")
	} else {
		return errors.New("not gonna happen")
	}

	return nil
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func parseMarkdown(data []byte) ast.Node {
	p := parser.New()
	return p.Parse(data)
}
