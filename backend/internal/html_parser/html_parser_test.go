package html_parser

import (
	"testing"
)

func TestParseHTML_ReturnsHeaderTagsCount(t *testing.T) {
	html := `
	<html>
		<body>
			<h1>FirstHeading</h1>
			<h1>Second Heading</h1>
			<h2>Subheading</h2>
			<h2>Another Subheading</h2>
			<h3>Third Level Heading</h3>
			<h4>Fourth Level Heading</h4>
			<h5>First Fifth Level Heading</h5>
			<h5>Second Fifth Level Heading</h5>
			<h5>Third Fifth Level Heading</h5>
			<h6>Sixth Level Heading</h6>
		</body>
	</html>`

	result, err := ParseHTML(html)
	if err != nil {
		t.Errorf("expected no error for parse, got: %v", err)
	}

	if result.H1TagsCount != 2 {
		t.Errorf("expected count of h1 tags to be 2, got %d", result.H1TagsCount)
	}
	if result.H2TagsCount != 2 {
		t.Errorf("expected count of h2 tags to be 2, got %d", result.H2TagsCount)
	}
	if result.H3TagsCount != 1 {
		t.Errorf("expected count of h3 tags to be 1, got %d", result.H3TagsCount)
	}
	if result.H4TagsCount != 1 {
		t.Errorf("expected count of h4 tags to be 1, got %d", result.H4TagsCount)
	}
	if result.H5TagsCount != 3 {
		t.Errorf("expected count of h5 tags to be 3, got %d", result.H5TagsCount)
	}
	if result.H6TagsCount != 1 {
		t.Errorf("expected count of h6 tags to be 1, got %d", result.H6TagsCount)
	}
}

func TestParseHTML_ReturnsPageTitle(t *testing.T) {
	html := `
	<!DOCTYPE html>
		<head>
			<title>Page Title</title>
		<body>
			<h1>FirstHeading</h1>
		</body>
	</html>`

	result, err := ParseHTML(html)
	if err != nil {
		t.Errorf("expected no error for parse, got: %v", err)
	}

	if result.Title != "Page Title" {
		t.Errorf("expected html version to be 'Page Title', got %s", result.Title)
	}
}

func TestParseHTML_ReturnsHTML5Version(t *testing.T) {
	html := `
	<!DOCTYPE html>
		<body>
			<h1>FirstHeading</h1>
		</body>
	</html>`

	result, err := ParseHTML(html)
	if err != nil {
		t.Errorf("expected no error for parse, got: %v", err)
	}

	if result.HTMLVersion != "5" {
		t.Errorf("expected html version to be 5, got %s", result.HTMLVersion)
	}
}
func TestParseHTML_ReturnsHTML4Version(t *testing.T) {
	html := `
	<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN">
		<body>
			<h1>FirstHeading</h1>
		</body>
	</html>`

	result, err := ParseHTML(html)
	if err != nil {
		t.Errorf("expected no error for parse, got: %v", err)
	}

	if result.HTMLVersion != "4.01" {
		t.Errorf("expected html version to be 4.01, got %s", result.HTMLVersion)
	}
}
