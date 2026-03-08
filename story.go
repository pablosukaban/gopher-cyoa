package gophercyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type handler struct {
	story  Story
	tmpl   *template.Template
	pathFn func(r *http.Request) string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	return path[1:]
}

func NewHandler(s Story, opts ...HandlerOptions) http.Handler {
	h := handler{s, tmpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if currentStory, ok := h.story[path]; ok {
		err := h.tmpl.Execute(w, currentStory)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Somsing wrong", http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "No such chapter", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story

	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
