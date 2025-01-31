package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURLSetXMLMarshal(t *testing.T) {

	urlset := sitemap.NewURLSet(
		*sitemap.NewURL("https://example.com"),
		*sitemap.NewURL("https://example.com/two").SetLastmodification("2024-01-02"),
	)

	data, err := xml.Marshal(urlset)
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://example.com</loc></url><url><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></url></urlset>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

}

func TestURLSetXMLUnmarshal(t *testing.T) {

	urlset := sitemap.NewURLSet()

	err := xml.Unmarshal([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://example.com</loc></url><url><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></url></urlset>"), urlset)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(urlset.URL) != 2 {
		t.Fatalf("Invalid result: \n%#v\n", urlset)
	}
}
