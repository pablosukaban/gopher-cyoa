package gophercyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []Option `json:"options,omitempty"`
}

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

const StoryTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Story}}<p>{{ . }}</p>{{end}}
		<p>Where now?</p>
		<ol>
			{{range .Options}}
			<li>
				<a href="{{.Arc}}">
					{{.Text}}
				</a>
			</li>
			{{end}}
		</ol>
	</body>
</html>`

func NewHandler(s Story) http.Handler {
	return MyHandler{stories: s}
}

type MyHandler struct {
	stories Story
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	storyName := path[1:]

	t, err := template.New("").Parse(StoryTemplate)
	if err != nil {
		log.Fatal(err)
	}

	if currentStory, ok := h.stories[storyName]; ok {
		err := t.Execute(w, currentStory)
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
