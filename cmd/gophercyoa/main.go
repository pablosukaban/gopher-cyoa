package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	gophercyoa "github.com/pablosukaban/gopher-cyoa"
)

func main() {
	fileName := flag.String("file", "gopher.json", "file name")
	port := flag.Int("port", 8080, "port")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := gophercyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	newTmpl := template.Must(template.New("").Parse(gophercyoa.StoryHandlerTmpl))

	handler := gophercyoa.NewHandler(story,
		gophercyoa.WithTemplate(newTmpl),
		gophercyoa.WithPathFn(pathFunc),
	)

	mux := http.NewServeMux()
	mux.Handle("/story/", handler)

	log.Printf("Listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}
