package report

import (
	// "errors"
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

	parseMarkdown(file)
	// if _, ok := root.(*ast.Document); ok {
	// 	fmt.Println("found the root")
	// } else {
	// 	return errors.New("not gonna happen")
	// }

	return nil
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func parseMarkdown(data []byte) {
	renderer := NewRenderer()
	extensions := parser.CommonExtensions | parser.Mmark
	parser := parser.NewWithExtensions(extensions)
	md := parser.Parse(data)
	ast.WalkFunc(md, renderer.RenderNode)
}

type Renderer struct {
	doc *Doc
}

func NewRenderer() *Renderer {
	return &Renderer{new(Doc)}
}

// headings, bulleted lists, text, preformatted text, image, background image
// iframe(?), link, html(?), caption (figurecaption), video
func (r *Renderer) RenderNode(node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Document:
		// do nothing
	case *ast.Paragraph:
		// do nothing
	case *ast.Text:
		r.text(node)
	case *ast.Heading:
		r.heading(node, entering)
	case *ast.HorizontalRule:
		// new section
	default:
		fmt.Println(fmt.Sprintf("Not implemented %T", node))
	}
	return ast.GoToNext
}

func (r *Renderer) text(node *ast.Text) {
	fmt.Println("text ", string(node.Literal))
}

func (r *Renderer) heading(node *ast.Heading, entering bool) {
	if entering {
		fmt.Println("heading ", string(node.Literal))
	}
}
