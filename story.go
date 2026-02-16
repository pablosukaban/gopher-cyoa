package gophercyoa

import (
	"encoding/json"
	"io"
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

func JsonStory(r io.Reader) (Story, error) {
	var story Story

	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
