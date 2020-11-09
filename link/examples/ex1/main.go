package main

import (
	"fmt"
	"strings"

	"github.com/gophercises/link"
)

var exampleHTML string = `<html>
<body>
  <h1>Hello!</h1>
	<a href="/other-page">wowww<span>A link to another page</span</a>
	<a href="/simple-page">A link to sample page</a>
</body>
</html>`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
