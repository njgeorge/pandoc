# pandocd

`pandocd` is a simple utility that converts markdown files (defined by extension
`.md`) to `html` using a `pandoc` backend. It assumes that `pandoc` is installed
and is *visible* in $PATH. The html output is refreshed every time that the
source markdown content is saved.

```
$> pandocd
  -debug
    	specity to enable debug mode
  -dst string
    	path to destination files (defaults to src if unspecified)
  -src string
    	path to source files (req.)
```

An output `.html` file is created for every `.md` file found in the folder
specified using the `-src` command line argument. If the `-dst` switch is not
specified, the output html files are created in the same folder as `-src`.

Using this you can use your favorite text editor (read vi) and focus purely on
content while the daemon in the background continuously keeps generating an
equivalent html that you can view in a browser. 

At this time the html file needs to be manually refreshed. Although an
auot-update feature is in the works.
