package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestIndex(t *testing.T) {

	i1 := sitemap.NewIndex(
		*sitemap.NewURL(sitemap.NewLocation("https://example.com/")).SetLastmodification(sitemap.NewLastModification("2024-01-01")),
		*sitemap.NewURL(sitemap.NewLocation("https://example.com/one")).SetLastmodification(sitemap.NewLastModification("2024-01-02")),
		*sitemap.NewURL(sitemap.NewLocation("https://example.com/two")).SetLastmodification(sitemap.NewLastModification("2024-01-03")),
	)

	data, err := i1.ToXML()
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><sitemap><loc>https://example.com/</loc><lastmod>2024-01-01</lastmod></sitemap><sitemap><loc>https://example.com/one</loc><lastmod>2024-01-02</lastmod></sitemap><sitemap><loc>https://example.com/two</loc><lastmod>2024-01-03</lastmod></sitemap></sitemapindex>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

	i2 := new(sitemap.Index)

	err = xml.Unmarshal(data, i2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(i1.Sitemaps) != len(i2.Sitemaps) {
		t.Fatalf("Invalid result:\n%#v\n%#v\n", *i1, *i2)
	}

}

func TestIndexToJSON(t *testing.T) {

	i := sitemap.NewIndex(
		*sitemap.NewURL(sitemap.NewLocation("https://example.com")),
		*sitemap.NewURL(sitemap.NewLocation("https://example.com/two")).SetLastmodification(sitemap.NewLastModification("2024-01-02")),
	)

	buf, err := i.ToJSON()
	if err != nil {
		t.Fatalf("Failed to marshal to TXT: %s\n", err)
	}

	if !bytes.Equal(buf, []byte(`{"sitemapindex":{"sitemap":[{"loc":"https://example.com"},{"loc":"https://example.com/two","lastmod":"2024-01-02"}]}}`)) {
		t.Fatalf("Invalid result: \n%s\n", buf)
	}
}
