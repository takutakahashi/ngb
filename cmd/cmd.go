package main

import (
	"flag"
	"log"
	"net/url"
	"os"

	"github.com/takutakahashi/ngb/pkg/load"
	"github.com/takutakahashi/ngb/pkg/result"
	"github.com/takutakahashi/ngb/pkg/types/config"
)

func main() {
	br, err := load.Start(parseArgument())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(result.Analyze(br))
}

func parseArgument() (*url.URL, int, int, config.ScriptConfig) {
	urlString := flag.String("url", "", "target URL")
	concurrency := flag.Int("c", 1, "Concurrency. default 1")
	num := flag.Int("n", 1, "Number of request per client. default 1")
	prerequest := flag.String("prerequest", "", "script path executing at pre request")
	flag.Parse()
	targetURL, err := url.Parse(*urlString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	sc := config.ScriptConfig{
		PreRequest: config.Script{Path: *prerequest},
	}
	return targetURL, *concurrency, *num, sc

}
