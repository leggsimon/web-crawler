package html_parser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ParseHTMLResult struct {
	H1TagsCount int
	H2TagsCount int
	H3TagsCount int
	H4TagsCount int
	H5TagsCount int
	H6TagsCount int
}

func ParseHTML(htmlContent string) (ParseHTMLResult, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	result := ParseHTMLResult{}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode {
			switch n.DataAtom {
			case atom.H1:
				result.H1TagsCount++
			case atom.H2:
				result.H2TagsCount++
			case atom.H3:
				result.H3TagsCount++
			case atom.H4:
				result.H4TagsCount++
			case atom.H5:
				result.H5TagsCount++
			case atom.H6:
				result.H6TagsCount++
			}
		}
	}
	return result, nil
}
