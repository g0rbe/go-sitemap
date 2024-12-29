package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURLSet(t *testing.T) {

	u1 := sitemap.NewURLSet(
		[]*sitemap.URL{
			sitemap.NewURL(sitemap.NewLoc("https://example.com"), nil, nil, nil),
			sitemap.NewURL(sitemap.NewLoc("https://example.com/two"), sitemap.NewLastMod("2024-01-02"), nil, nil),
		})

	data, err := u1.ToXML()
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://example.com</loc></url><url><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></url></urlset>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

	u2 := new(sitemap.URLSet)

	err = xml.Unmarshal(data, u2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(u1.URLs) != len(u2.URLs) {
		t.Fatalf("Invalid result: \n%#v\n%#v\n", *u1, *u2)
	}
}

func TestFetchURLSet(t *testing.T) {

	u, err := sitemap.FetchURLSet("https://gorbe.io/en/sitemap.xml")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	if len(u.URLs) == 0 {
		t.Fatalf("Invalid result: zero length\n")
	}
}

func ExampleFetchURLSet() {

	u, err := sitemap.FetchURLSet("https://gorbe.io/en/sitemap.xml")
	if err != nil {
		// Handle error
	}

	fmt.Printf("Length is %d\n", len(u.URLs))
}
