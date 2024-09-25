package app

import (
	"embed"
	"html/template"
	"io/fs"
	"path"
)

type templates struct {
	templateFuncs template.FuncMap
	baseTemplate  *template.Template
	TrackerTmpl   *template.Template
	ErrorTmpl     *template.Template
}

//go:embed templates
var embeddedFs embed.FS
var webFs fs.FS = embeddedFs

var (
	templatesDir  = "templates"
	componentsDir = path.Join(templatesDir, "components")
	pagesDir      = path.Join(templatesDir, "pages")
)

func (t *templates) parseTemplates() {
	t.templateFuncs = template.FuncMap{}

	t.baseTemplate = template.Must(template.New("base.html").
		Funcs(t.templateFuncs).
		ParseFS(webFs, path.Join(templatesDir, "layouts", "base.html")),
	)

	t.TrackerTmpl = t.pageTmpl("tracker.html")
	t.ErrorTmpl = t.pageTmpl("error.html")
}

func (t *templates) pageTmpl(fileName string) *template.Template {
	return template.Must(
		template.Must(t.baseTemplate.Clone()).ParseFS(
			webFs,
			path.Join(pagesDir, fileName),
			path.Join(componentsDir, "*.html")))
}
