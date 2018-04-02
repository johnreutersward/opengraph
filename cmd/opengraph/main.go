package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/johnreutersward/opengraph"
)

const (
	version = "0.0.2"
)

func printUsage() {
	flag.Usage()
	os.Exit(0)
}

func printVersion() {
	fmt.Fprintf(os.Stdout, "opengraph v%s\n", version)
	os.Exit(0)
}

func main() {
	var (
		showHelp    = flag.Bool("help", false, "Show usage help")
		showVersion = flag.Bool("version", false, "Show version")
		prefix      = flag.String("prefix", "og", "Prefix")
		outputJson  = flag.Bool("json", false, "Output in JSON")
	)

	flag.Parse()

	if *showVersion {
		printVersion()
	}

	if flag.NArg() == 0 || *showHelp {
		printUsage()
	}

	args := flag.Args()
	url := args[0]

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	md, err := opengraph.ExtractPrefix(res.Body, *prefix)
	if err != nil {
		log.Fatal(err)
	}

	if *outputJson {
		data, err := json.MarshalIndent(&md, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data)
		fmt.Fprintf(os.Stdout, "\n")
	} else {
		for i := range md {
			fmt.Fprintf(os.Stdout, "%s: %s\n", md[i].Property, md[i].Content)
		}
	}

	os.Exit(0)
}
