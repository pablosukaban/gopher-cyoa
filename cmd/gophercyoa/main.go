package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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

	handler := gophercyoa.NewHandler(story)

	log.Printf("Listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
