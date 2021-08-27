package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			fmt.Println(path)
			fmt.Println(dest)
			http.Redirect(w, r, dest, http.StatusPermanentRedirect)
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
func YAMLHandler(yamlFile string, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := yamlToPaths(yamlFile)
	if err != nil {
		return nil, err
	}
	pathsUrl, err := makeMap(paths)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsUrl, fallback), nil
}

func JSONHandler(jsonFile string, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := jsonToPaths(jsonFile)
	if err != nil {
		return nil, err
	}
	pathsUrl, err := makeMap(paths)
	if err != nil {
		return nil, err
	}
	fmt.Println(pathsUrl)
	return MapHandler(pathsUrl, fallback), nil
}

func makeMap(paths []pathUrl) (map[string]string, error) {
	var ret = make(map[string]string, len(paths))
	for _, pu := range paths {
		ret[pu.Path] = pu.URL
	}
	return ret, nil
}

func yamlToPaths(yamlFile string) ([]pathUrl, error) {
	f, err := os.Open(yamlFile)
	if err != nil {
		return nil, err
	}
	var paths []pathUrl
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func jsonToPaths(jsonFile string) ([]pathUrl, error) {
	f, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	var paths []pathUrl
	dec := json.NewDecoder(f)
	err = dec.Decode(&paths)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

type pathUrl struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url"`
}
