package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestIndex(t *testing.T) {

	i1 := sitemap.NewIndex(
		sitemap.NewURL(sitemap.NewLocation("https://example.com/"), sitemap.NewLastMod("2024-01-01"), nil, nil),
		sitemap.NewURL(sitemap.NewLocation("https://example.com/one"), sitemap.NewLastMod("2024-01-02"), nil, nil),
		sitemap.NewURL(sitemap.NewLocation("https://example.com/two"), sitemap.NewLastMod("2024-01-03"), nil, nil),
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
