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

	v, err := s.ToXMLIndent()
	if err != nil {
		t.Fatalf("Failed to convert to XML: %s\n", err)
	}

	t.Logf("\n%s\n", v)
}

func TestRemoveURL(t *testing.T) {

	sm := sitemap.New()
	sm.SetURL(sitemap.NewURL("https://example.com/0"))
	sm.SetURL(sitemap.NewURL("https://example.com/1"))
	sm.SetURL(sitemap.NewURL("https://example.com/2"))
	sm.SetURL(sitemap.NewURL("https://example.com/3"))

	sm.RemoveURL("https://example.com/2")
	u2 := sm.GetURL("https://example.com/2")
	if u2 != nil {
		t.Fatalf("Failed to remove the /2 url\n")
	}

	sm.RemoveURL("https://example.com/0")
	u0 := sm.GetURL("https://example.com/0")
	if u0 != nil {
		t.Fatalf("Failed to remove the /0 url\n")
	}

	sm.RemoveURL("https://example.com/3")
	u3 := sm.GetURL("https://example.com/3")
	if u3 != nil {
		t.Fatalf("Failed to remove the /3 url\n")
	}

	if sm.Size() != 1 {
		t.Fatalf("Invalid size: %d\n", sm.Size())
	}

}
func TestCrawler(t *testing.T) {
	s, err := sitemap.CrawlHostname("example.com", nil)
	if err != nil {
		t.Fatalf("Fail: %s\n", err)
	}

	if s.Size() == 0 {
		t.Fatalf("Empty result")
	}
}
