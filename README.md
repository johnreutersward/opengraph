# opengraph

[![Build Status](https://travis-ci.org/johnreutersward/opengraph.svg?branch=master)](https://travis-ci.org/johnreutersward/opengraph)
[![Go Reference](https://pkg.go.dev/badge/github.com/johnreutersward/opengraph.svg)](https://pkg.go.dev/github.com/johnreutersward/opengraph)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnreutersward/opengraph)](https://goreportcard.com/report/github.com/johnreutersward/opengraph)

![ogp](ogp.png?raw=true "ogp")

> opengraph is a Go library and command-line tool for extracting [Open Graph](https://ogp.me/ "Open Graph protocol") metadata from HTML documents.

## Library

### Install

```go
import "github.com/johnreutersward/opengraph"
```

### Usage

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

## Command-line tool

### Install

Binary releases: 

https://github.com/johnreutersward/opengraph/releases

Or build from source:

```
$ go get github.com/johnreutersward/opengraph/cmd/opengraph
```

### Usage

```
$ opengraph http://www.imdb.com/title/tt0118715/
type: video.movie
title: The Big Lebowski (1998)
site_name: IMDb
...
```

Output in JSON:

```
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
