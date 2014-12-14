package main

import . "github.com/puerkitobio/goquery"
import "github.com/tj/docopt"
import "strings"
import "fmt"
import "os"

var usage = `
  Usage: 
    pick [options] <selector>...
    pick -v | --version
    pick -h | --help

  Options:
    -t, --text          output inner text for each element
    -v, --version       show version information
    -h, --help          show help information

  Examples:

    # pick <title> and output it's text.
    echo "<title>foo</title>" | pick title --text

    # pick <script>s and output sources.
    curl --silent https://github.com | pick script :src

    # pick all <span>'s in <a>.
    curl --silent https://github.com | pick a | pick span
    curl --silent https://github.com | pick "a span"

    # pick all [src=*].
    curl --silent https://github.com | pick :href

    # pick all spans and anchors.
    curl --silent https://github.com | pick a span
`

func main() {
	args, err := docopt.Parse(usage, nil, true, "0.0.1", false)
	check(err)

	var selectors []string
	var attrs []string
	text := args["--text"].(bool)

	for _, s := range args["<selector>"].([]string) {
		if ':' != s[0] {
			selectors = append(selectors, s)
		} else {
			attrs = append(attrs, s[1:])
		}
	}

	doc, err := NewDocumentFromReader(os.Stdin)
	check(err)

	if len(selectors) == 0 {
		selectors = append(selectors, "*")
	}

	selection := doc.Find(strings.Join(selectors, ", "))

	selection.Each(func(_ int, s *Selection) {
		if text {
			fmt.Printf("%s\n", s.Text())
		}

		if len(attrs) > 0 {
			for _, attr := range attrs {
				if v, ok := s.Attr(attr); ok {
					fmt.Printf("%s\n", v)
				}
			}
		}

		if !text && 0 == len(attrs) {
			v, _ := s.Html()
			fmt.Printf("%s\n", v)
		}
	})
}

func check(err error) {
	if err != nil {
		fmt.Printf("pick: %s\n", err.Error())
		os.Exit(1)
	}
}
