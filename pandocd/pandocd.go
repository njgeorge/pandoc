package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var htmltop string = `
<!DOCTYPE html>
<html>
<head>
<style>
div {
    width:700px;
    background-color:#333333;
    color:lightgray;
    margin:0 auto;
    padding:10px;
    font-size:20px;
}
</style>
</head>
<body bgcolor="#000000">
<div>
`

var htmlbot string = `
</div>
</body>
</html>
`

func Convert(path string) {
	out, err := exec.Command("pandoc", "-f", "markdown", "-t", "html", path).Output()
	if err != nil {
		log.Fatal(err)
	}

	_, fname := filepath.Split(path)
	opath := dst + "/" + strings.TrimSuffix(fname, ".md") + ".html"
	log.Println(opath)

	file, err := os.Create(opath)
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(htmltop))
	file.Write(out)
	file.Write([]byte(htmlbot))

	file.Close()
	log.Println("Wrote", opath)
}

func FilterEvents(watcher *fsnotify.Watcher) {
	for {
		select {
		case evt := <-watcher.Events:
			if strings.HasSuffix(evt.Name, ".md") &&
				(evt.Op&fsnotify.Write != 0 || evt.Op&fsnotify.Create != 0) {
				Convert(evt.Name)
			}
		case err := <-watcher.Errors:
			log.Fatal(err)
		}
	}
}

var src, dst string

func main() {
	var dbg bool

	flag.StringVar(&src, "src", "", "path to source files (req.)")
	flag.StringVar(&dst, "dst", "", "path to destination files (defaults to src if unspecified)")
	flag.BoolVar(&dbg, "debug", false, "specity to enable debug mode")

	flag.Parse()

	log.SetFlags(0)
	if dbg {
		log.SetFlags(log.Lshortfile)
	}

	if src == "" {
		flag.PrintDefaults()
		log.Fatal("-src is a required argument.")
	}

	if dst == "" {
		dst = src
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(src)
	if err != nil {
		log.Fatal(err)
	}

	FilterEvents(watcher)
}
