// package opengraph is a library for extracting OpenGraph meta-data from an html document.
// See http://ogp.me/ for more information about the OpenGraph project.
package main

import (
	//"errors"
	"fmt"
	"io"
	"strings"

	"code.google.com/p/go.net/html"
)

type OpenGraph struct {
	Title *string
	Type  *string
	Image *string
	Url   *string
}

type ogAttr struct {
	property string
	content  string
}

// Extract extracts the OpenGraph data from a html document.
// The input is assumed to be UTF-8 encoded.
func Extract(r io.Reader) (*OpenGraph, error) {

	tags, err := ogAttrs(r)

	if err != nil {
		return nil, err
	}

	fmt.Println(tags)

	return nil, nil

}

// ogAttrs extracts the OpenGraph attributes from meta tags.
func ogAttrs(r io.Reader) ([]ogAttr, error) {
	var tags []ogAttr
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				return tags, nil
			}
			return nil, z.Err()
		}

		t := z.Token()

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

			if strings.HasPrefix(prop, "og:") {
				tags = append(tags, ogAttr{prop, cont})
			}
		}
	}

	return tags, nil
}

// test, remove later
func main() {
	//s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	s := `<meta property='' content="http://ia.media-imdb.com/images/M/MV5BNjc1NzYwODEyMV5BMl5BanBnXkFtZTcwNTcxMzU1MQ@@._V1_SY1200_CR126,0,630,1200_AL_.jpg" />
	<meta property='og:type' content="video.tv_show" />
    <meta property='fb:app_id' content='115109575169727' />
    <meta property='og:title' content="The Wire (TV Series 2002–2008)" />
    <meta property='og:site_name' content='IMDb' />`
	rdr := strings.NewReader(s)
	(Extract(rdr))
}
