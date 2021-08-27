package main

import (
	"flag"
	"fmt"
	"net/http"
	"urlshortMain/urlshort"
)

func main() {
	yamlFile := flag.String("yaml", "", "Yaml file for the YamlHandler")
	jsonFile := flag.String("json", "", "Json file for the JsonHanlder")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello User, You have reached the end of the tunnel. This the end of the internet as you know it")
	})
	pathsUrl := map[string]string{
		"/g": "https://google.com",
		"/y": "https://youtube.com",
	}
	h := urlshort.MapHandler(pathsUrl, mux)
	if *yamlFile != "" {
		var err error
		h, err = urlshort.YAMLHandler(*yamlFile, h)
		if err != nil {
			panic(err)
		}
	}
	if *jsonFile != "" {
		var err error
		h, err = urlshort.JSONHandler(*jsonFile, h)
		if err != nil {
			panic(err)
		}
	}

	http.ListenAndServe(":3000", h)

}
