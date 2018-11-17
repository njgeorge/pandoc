package main

import (
	"flag"
	"log"
)

func main() {
	var src, dst string
	var dbg bool

	flag.StringVar(&src, "src", "", "path to source files")
	flag.StringVar(&dst, "dst", "", "path to destination files")
	flag.BoolVar(&dbg, "debug", false, "specity to enable debug mode")

	flag.Parse()

	log.SetFlags(0)
	if dbg {
		log.SetFlags(log.Lshortfile)
	}

	if src == "" || dst == "" {
		flag.PrintDefaults()
		log.Fatal("src and dst are required.")
	}
}
