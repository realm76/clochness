package components

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Page represents a base page layout component
type Page struct {
	Title string
	Child Node
}

// NewPage creates a new Page component
func NewPage(title string, child Node) *Page {
	return &Page{
		Title: title,
		Child: child,
	}
}

func (p *Page) Render() Node {
	return HTML5(HTML5Props{
		Title:    p.Title,
		Language: "en",
		Head: []Node{
			Link(Href("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css")),
		},
		Body: []Node{
			Div(),

			Script(Src("https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js")),
		},
	})

}
