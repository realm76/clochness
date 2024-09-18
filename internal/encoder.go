package encoder

import (
	"go.uber.org/zap"
	"html/template"
	"io"
	"net/http"
)

type Encoder struct {
	templates *template.Template
	logger    *zap.SugaredLogger
}

type Component interface {
	Render(w io.Writer) error
}

type HtmlEncoder interface {
	EncodeHTML(w http.ResponseWriter, r *http.Request, status int, contents string)
	EncodeComponent(w http.ResponseWriter, r *http.Request, status int, component Component)
	EncodeError(w http.ResponseWriter, err error)
}

func NewEncoder(logger *zap.SugaredLogger, templates *template.Template) *Encoder {
	if templates == nil {
		panic("nil templates")
	}

	return &Encoder{
		templates: templates,
		logger:    logger,
	}
}

func (e *Encoder) EncodeComponent(w http.ResponseWriter, r *http.Request, status int, component Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	err := component.Render(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (e *Encoder) EncodeHTML(w http.ResponseWriter, r *http.Request, status int, contents string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	_, err := w.Write([]byte(contents))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (e *Encoder) EncodeError(w http.ResponseWriter, errToEncode error) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(500)

	e.logger.Errorln(errToEncode.Error())

	if err := e.templates.ExecuteTemplate(w, "error", nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
