package urlshort

import (
	"fmt"
	"net/http"

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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsUrl, err := makeMap(yml)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsUrl, fallback), nil
}

func makeMap(yml []byte) (map[string]string, error) {
	var paths []pathUrl
	err := yaml.Unmarshal(yml, &paths)
	if err != nil {
		return nil, err
	}
	fmt.Println(paths)
	var ret = make(map[string]string, len(paths))
	for _, pu := range paths {
		ret[pu.Path] = pu.URL
	}
	return ret, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
