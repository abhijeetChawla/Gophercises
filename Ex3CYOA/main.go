package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	mux := http.NewServeMux()
	arcs, err := parseArcs()
	if err != nil {
		panic(err)
	}
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "intro"
		} else {
			path = path[1:]
		}
		data := arcs[path]
		tmpl := template.Must(template.ParseFiles("tempalte.html"))
		err := tmpl.Execute(rw, data)
		if err != nil {
			log.Fatalf("template execution: %s", err)
		}
	})

	http.ListenAndServe(":3000", mux)
}

func parseArcs() (map[string]arc, error) {
	b, err := os.ReadFile("gopher.json")
	if err != nil {
		return nil, err
	}
	var arcs map[string]arc
	err = json.Unmarshal(b, &arcs)
	if err != nil {
		return nil, err
	}
	return arcs, nil
}

type arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}

type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
