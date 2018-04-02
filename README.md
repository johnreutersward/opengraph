# [![Open Graph protocol](http://imgur.com/pqFdEuo.png?1)](http://ogp.me/) opengraph

opengraph is a Go library and command-line tool for extracting Open Graph metadata from HTML documents.

**Documentation:** <http://godoc.org/github.com/johnreutersward/opengraph>  
**Open Graph protocol:** <http://ogp.me/>  
**Build Status:** [![travis-ci status](https://api.travis-ci.org/johnreutersward/opengraph.png)](https://travis-ci.org/johnreutersward/opengraph)  

## Usage

```go
import "github.com/johnreutersward/opengraph"
```

To extract Open Graph metadata from a movie on IMDb (sans error handling)

```go
res, _ := http.Get("http://www.imdb.com/title/tt0118715/")
md, _ := opengraph.Extract(res.Body)
for i := range md {
	fmt.Printf("%s = %s\n", md[i].Property, md[i].Content)
}
```

Which will output

```
url = http://www.imdb.com/title/tt0118715/
type = video.movie
title = The Big Lebowski (1998)
site_name = IMDb
description = Directed by Joel Coen, Ethan Coen.  With Jeff Bridges ...
...
```

## Command-line usage

Install

```bash
$ go get github.com/johnreutersward/opengraph/cmd/opengraph
```

Run

```bash
$ opengraph http://www.imdb.com/title/tt0118715/
type: video.movie
title: The Big Lebowski (1998)
site_name: IMDb
...
```

Output in JSON

```bash
$ opengraph -json http://www.imdb.com/title/tt0118715/
[
  {
    "Property": "type",
    "Content": "video.movie",
    "Prefix": "og"
  },
  {
    "Property": "title",
    "Content": "The Big Lebowski (1998)",
    "Prefix": "og"
  },
  {
    "Property": "site_name",
    "Content": "IMDb",
    "Prefix": "og"
  },
  ...
]
```

## License

MIT
