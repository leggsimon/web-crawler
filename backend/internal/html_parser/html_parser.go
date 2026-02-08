package html_parser

import (
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ParseHTMLResult struct {
	HTMLVersion string
	Title       string
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
		if n.Type == html.DoctypeNode {
			result.HTMLVersion = extractHTMLVersion(n)
		}

		if n.Type == html.ElementNode {
			switch n.DataAtom {
			case atom.Title:
				result.Title = n.FirstChild.Data
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

func extractHTMLVersion(n *html.Node) string {
	if strings.ToLower(n.Data) == "html" {
		if n.Attr == nil {
			return "5"
		}

		for _, a := range n.Attr {
			val := strings.ToLower(a.Val)

			// TODO: Theres definitely a more comprehensive way to get all but for now this will do.
			switch {
			case strings.Contains(val, "html 4.01"):
				return "4.01"
			case strings.Contains(val, "html 4.00"):
				return "4.00"
			}
		}
	}
	return "Unknown"
}
