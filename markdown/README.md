Processor Wrapper For [blackfriday](https://github.com/russross/blackfriday)
==========

## Usage

The example code below will run a http web server on port `8000`.
It will then serve files from a local directory with the name `markdown`.
For example, populate a markdown file in the markdown directory `./markdown/test.html.md`.

> note: I use `.html.md` extension for mime type purposes.

Visit this file in the browser as `localhost:8000/test.html` (stripped `.md` prefix).
Notice it renders the file to browser in html.

## Example Code
```go
package main

import (
	"net/http"

	"github.com/GeorgeMac/go-assets"
	"github.com/GeorgeMac/go-assets-processors/markdown"
	"github.com/GeorgeMac/go-assets/cache"
)

func main() {
	md := markdown.New()
	c := cache.New(cache.Proc(md), cache.Match(".md"))
	a := assets.New("/", assets.Dir("markdown"), assets.SetCache(c))
	http.Handle(a.Pattern(), a)
	http.ListenAndServe(":8000", nil)
}
```
