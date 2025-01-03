package sitemap_test

import (
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestFetch(t *testing.T) {

	s, err := sitemap.Fetch("https://gorbe.io/sitemap.xml")
	if err != nil {
		t.Fatalf("Failed to fetch: %s\n", err)
	}

	v, err := s.ToXML()
	if err != nil {
		t.Fatalf("Failed to convert to XML: %s\n", err)
	}

	t.Logf("\n%s\n", v)
}
