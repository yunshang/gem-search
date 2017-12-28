package main

import (
  "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-wordwrap"
	gopen "github.com/petermbenjamin/go-open"
	"github.com/tj/docopt"
)

type Response []struct {
    Name     string
    Downloads int64
    Authors   string
    Info string
    Source_code_uri  string
}
// Version is the package version
var Version = "0.0.1"

// Usage is the package usage information
const Usage = `
  Usage:
    gem-search <query>... [--top] [--count n] [--open]
    gem-search -h | --help
    gem-search --version

  Options:
    -n, --count n    number of results [default: 5]
    -o, --open       open godoc.org search results in default browser
    -h, --help       output help information
    -v, --version    output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
  }

	n, err := strconv.ParseInt(args["--count"].(string), 10, 32)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	query := strings.Join(args["<query>"].([]string), " ")
	res, err := http.Get("https://rubygems.org/api/v1/search.json?query=" + url.QueryEscape(query))
	if err != nil {
		log.Fatalf("request failed: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("request error: %s", http.StatusText(res.StatusCode))
	}

	var body Response
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		log.Fatalf("error parsing response: %s", err)
  }

	if open := args["--open"].(bool); open {
		gopen.Open("https://rubygems.org/search?/?query=" + url.QueryEscape(query))
		os.Exit(0)
	}

	body = body[:n]

	println()
	for _, pkg := range body {
		fmt.Printf("  \033[1m%s\033[m\n", pkg.Name)
		fmt.Printf("  %s\n", pkg.Source_code_uri)
		fmt.Printf("  %s\n", description(pkg.Info))
		fmt.Printf("\n")
	}
	println()
}

func description(s string) string {
	if s == "" {
		return "no info"
	}

	s = wordwrap.WrapString(s, 60)
	s = strings.Replace(s, "\n", "\n  ", -1)
	return s
}
