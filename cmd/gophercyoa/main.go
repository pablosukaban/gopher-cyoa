package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	gophercyoa "github.com/pablosukaban/gopher-cyoa"
)

type MyHandler struct {
	stories gophercyoa.Story
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
	fileName := flag.String("file", "gopher.json", "file name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := gophercyoa.JsonStory(file)

	mux := http.NewServeMux()
	mh := MyHandler{story}

	mux.Handle("/", mh)

	log.Print("Listening...")
	http.ListenAndServe(":8080", mux)
}
