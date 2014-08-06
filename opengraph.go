// Package opengraph extracts Open Graph metadata from html documents.
// See http://ogp.me/ for more information about the Open Graph protocol.
//
// Usage:
// 	import "github.com/rojters/opengraph"
//
// To extract Open Graph metadata from a movie on IMDb:
//
// 	res, _ := http.Get("http://www.imdb.com/title/tt0118715/")
// 	og, _ := opengraph.Extract(res.Body)
// 	for _, md := range og {
// 		fmt.Printf("%s = %s\n", md.Property, md.Content)
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

	"code.google.com/p/go.net/html"
)

type MetaData struct {
	Property string // Porperty attribute without namespace prefix.
	Content  string // See http://ogp.me/#data_types for a list of content attribute types.
}

// By default Extract will only return metadata in the Open Graph namespace.
// This variable can be changed to get data from other namespaces.
// Ex: 'fb:' for Facebook or to get all metadata regardless of namespace, set it to the empty string.
var Namespace = "og:"

// Extract extracts Open Graph metadata from a html document.
// If no relevant metadata is found the result will be empty.
// The input is assumed to be UTF-8 encoded.
func Extract(doc io.Reader) ([]MetaData, error) {
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

		if t.Type == html.EndTagToken && t.Data == "head" {
			return tags, nil
		}

		if t.Type == html.SelfClosingTagToken && t.Data == "meta" {
			var prop, cont string
			for _, a := range t.Attr {
				switch a.Key {
				case "property":
					prop = a.Val
				case "content":
					cont = a.Val
				}
			}

			if strings.HasPrefix(prop, Namespace) && cont != "" {
				tags = append(tags, MetaData{prop[len(Namespace):], cont})
			}
		}
	}

	return tags, nil
}
