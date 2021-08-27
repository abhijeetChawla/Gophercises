package urlshort

import (
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
	paths, err := makePathsUrl(yamlFile)
	if err != nil {
		return nil, err
	}
	pathsUrl, err := makeMap(paths)
	if err != nil {
		return nil, err
	}
	return MapHandler(pathsUrl, fallback), nil
}

func makeMap(paths []pathUrl) (map[string]string, error) {
	var ret = make(map[string]string, len(paths))
	for _, pu := range paths {
		ret[pu.Path] = pu.URL
	}
	return ret, nil
}

func makePathsUrl(yml string) ([]pathUrl, error) {
	f, err := os.Open(yml)
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

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
