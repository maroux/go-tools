package main

import (
	"fmt"
	"os"
	"plugin"

	"golang.org/x/tools/godoc"
)

func registerPlugin(p *godoc.Presentation, pluginPath string) {
	pl, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load plugin at: %s: %v\n", pluginPath, err)
		os.Exit(1)
	}
	symbol, err := pl.Lookup("URLForSrc")
	if err == nil {
		if fn, ok := symbol.(func(src string) string); !ok {
			fmt.Fprintf(os.Stderr, "signature for URLForSrc in plugin: %s is not valid, expected: func(src string) string, found: %+v\n", pluginPath, symbol)
			os.Exit(1)
		} else {
			p.URLForSrc = fn
		}
	}
	symbol, err = pl.Lookup("URLForSrcPos")
	if err == nil {
		if fn, ok := symbol.(func(src string, line, low, high int) string); !ok {
			fmt.Fprintf(os.Stderr, "signature for URLForSrcPos in plugin: %s is not valid, expected: func(src string, line, low, high int) string, found: %+v\n", pluginPath, symbol)
			os.Exit(1)
		} else {
			p.URLForSrcPos = fn
		}
	}
	symbol, err = pl.Lookup("URLForSrcQuery")
	if err == nil {
		if fn, ok := symbol.(func(src, query string, line int) string); !ok {
			fmt.Fprintf(os.Stderr, "signature for URLForSrc in plugin: %s is not valid, expected: func(src, query string, line int) string, found: %+v\n", pluginPath, symbol)
			os.Exit(1)
		} else {
			p.URLForSrcQuery = fn
		}
	}
}
