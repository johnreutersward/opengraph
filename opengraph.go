// Package opengraph extracts Open Graph metadata from HTML documents.
// See http://ogp.me/ for more information about the Open Graph protocol.
//
// Usage:
// 	import "github.com/rojters/opengraph"
//
// To extract Open Graph metadata from a movie on IMDb (sans error handling):
//
// 	res, _ := http.Get("http://www.imdb.com/title/tt0118715/")
// 	md, _ := opengraph.Extract(res.Body)
// 	for i := range md {
// 		fmt.Printf("%s = %s\n", md[i].Property, md[i].Content)
// 	}
//
// Which will output:
//
// 	url = http://www.imdb.com/title/tt0118715/
// 	type = video.movie
// 	title = The Big Lebowski (1998)
// 	site_name = IMDb
// 	description = Directed by Joel Coen, Ethan Coen.  With Jeff Bridges ...
// 	...
//
package opengraph

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

const (
	defaultPrefix = "og"
)

type MetaData struct {
	Property string // Property attribute without prefix.
	Content  string // Content attribute. See http://ogp.me/#data_types for a list of content attribute types.
	Prefix   string
}

// Extract extracts Open Graph metadata from a HTML document.
// If no relevant metadata is found the result will be empty.
// The input is assumed to be UTF-8 encoded.
func Extract(doc io.Reader) ([]MetaData, error) {
	return ExtractPrefix(doc, defaultPrefix)
}

// Same as Extract but extracts metadata with a specific prefix, e.g. "fb" for Facebook.
// If prefix is empty all matching metadata is extracted.
func ExtractPrefix(doc io.Reader, prefix string) ([]MetaData, error) {
	var tags []MetaData
	z := html.NewTokenizer(doc)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				return tags, nil
			}
			return nil, z.Err()
		}

		t := z.Token()

		if t.Data == "head" && t.Type == html.EndTagToken {
			return tags, nil
		}

		if t.Data == "meta" {
			var prop, cont, name, tagPrefix string
			for _, a := range t.Attr {
				switch a.Key {
				case "property":
					prop = a.Val
				case "name":
					name = a.Val
				case "content":
					cont = a.Val
				}
			}

			if prop == "" {
				prop = name
			}

			if prop == "" || cont == "" {
				continue
			}

			if prefix != "" {
				if !strings.HasPrefix(prop, prefix+":") {
					continue
				}
				tagPrefix = prefix
			} else {
				idx := strings.Index(prop, ":")
				if idx == -1 {
					continue
				}
				tagPrefix = prop[:idx]
			}

			tags = append(tags, MetaData{prop[len(tagPrefix+":"):], cont, tagPrefix})
		}
	}

	return tags, nil
}
