package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshortMain/urlshort"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello User, You have reached the end of the tunnel. This the end of the internet as you know it")
	})
	pathsUrl := map[string]string{
		"/g": "https://google.com",
		"/y": "https://youtube.com",
	}
	h := urlshort.MapHandler(pathsUrl, mux)
	b, err := ioutil.ReadFile("url.yaml")
	if err != nil {
		panic(err)
	}
	ymalH, err := urlshort.YAMLHandler(b, h)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":3000", ymalH)

}
