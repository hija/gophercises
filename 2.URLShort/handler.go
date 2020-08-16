package urlshort

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Here we will extract the page title from the Request,
		// and call the provided handler 'fn'
		if redirectionurl, ok := pathsToUrls[r.RequestURI]; ok {
			http.Redirect(w, r, redirectionurl, 302)
		} else {
			// Not found
			fallback.ServeHTTP(w, r)
		}
	}

}

type urlShortFormat struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	urlShortformats := []urlShortFormat{}
	err := yaml.Unmarshal(yml, &urlShortformats)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	pathMap := buildMap(urlShortformats)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(yamlURLs []urlShortFormat) map[string]string {
	pathsToURLMap := make(map[string]string)

	for _, urlShortFormat := range yamlURLs {
		pathsToURLMap[urlShortFormat.Path] = urlShortFormat.URL
		fmt.Println(urlShortFormat.Path, "-->", urlShortFormat.URL)
	}

	return pathsToURLMap
}
