package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Option struct {
	Text string
	Arc  string
}

type Story struct {
	Title   string
	Story   []string
	Options []Option
}

type MyHandler struct {
	stories map[string]Story
}

const tpl = `
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

func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storyName := r.URL.Path[1:]

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	if currentStory, ok := mh.stories[storyName]; ok {
		t.Execute(w, currentStory)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	stories := make(map[string]Story)

	file, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &stories)

	mux := http.NewServeMux()
	mh := MyHandler{stories}

	mux.Handle("/", mh)

	log.Print("Listening...")
	http.ListenAndServe(":8080", mux)
}
