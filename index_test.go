package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestIndexXMLMarshal(t *testing.T) {

	index := sitemap.Index{}

	index = append(index, sitemap.URL{Location: "https://example.com/", LastModification: "2024-01-01"})
	index = append(index, sitemap.URL{Location: "https://example.com/one", LastModification: "2024-01-02"})
	index = append(index, sitemap.URL{Location: "https://example.com/two", LastModification: "2024-01-03"})

	data, err := xml.Marshal(index)
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><sitemap><loc>https://example.com/</loc><lastmod>2024-01-01</lastmod></sitemap><sitemap><loc>https://example.com/one</loc><lastmod>2024-01-02</lastmod></sitemap><sitemap><loc>https://example.com/two</loc><lastmod>2024-01-03</lastmod></sitemap></sitemapindex>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

}

func TestIndexXMLUnmarshal(t *testing.T) {

	index := sitemap.Index{}

	err := xml.Unmarshal([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><sitemap><loc>https://example.com/</loc><lastmod>2024-01-01</lastmod></sitemap><sitemap><loc>https://example.com/one</loc><lastmod>2024-01-02</lastmod></sitemap><sitemap><loc>https://example.com/two</loc><lastmod>2024-01-03</lastmod></sitemap></sitemapindex>"), &index)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(index) != 3 {
		t.Fatalf("Invalid result:\n%#v\n", index)
	}
}
