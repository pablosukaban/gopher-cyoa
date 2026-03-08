package gophercyoa

import (
	"html/template"
	"net/http"
)

type HandlerOptions func(h *handler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) {
		h.tmpl = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOptions {
	return func(h *handler) {
		h.pathFn = fn
	}
}
