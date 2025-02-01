package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestSitemapXMLMarshalURLSet(t *testing.T) {

	sm := sitemap.New()

	sm.URL = append(sm.URL, sitemap.URL{Location: "https://example.com"})
	sm.URL = append(sm.URL, sitemap.URL{Location: "https://example.com/two", LastModification: "2024-01-02"})

	data, err := xml.Marshal(sm)
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://example.com</loc></url><url><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></url></urlset>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

}

func TestSitemapXMLUnmarshalURLSet(t *testing.T) {

	sm := sitemap.New()

	err := xml.Unmarshal([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url><loc>https://example.com</loc></url><url><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></url></urlset>"), &sm)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(sm.URL) != 2 {
		t.Fatalf("Invalid result: \n%#v\n", sm)
	}
}

func TestSitemapXMLMarshalIndex(t *testing.T) {

	sm := sitemap.New()
	sm.SetIndex(true)

	sm.URL = append(sm.URL, sitemap.URL{Location: "https://example.com"})
	sm.URL = append(sm.URL, sitemap.URL{Location: "https://example.com/two", LastModification: "2024-01-02"})

	data, err := xml.Marshal(sm)
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	if !bytes.Equal(data, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><sitemap><loc>https://example.com</loc></sitemap><sitemap><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></sitemap></sitemapindex>")) {
		t.Fatalf("Invalid data: %s\n", data)
	}

}

func TestSitemapXMLUnmarshalIndex(t *testing.T) {

	sm := sitemap.New()

	err := xml.Unmarshal([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><sitemap><loc>https://example.com</loc></sitemap><sitemap><loc>https://example.com/two</loc><lastmod>2024-01-02</lastmod></sitemap></sitemapindex>"), &sm)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(sm.URL) != 2 {
		t.Fatalf("Invalid result: \n%#v\n", sm)
	}
}

func TestSitemapSetIndex(t *testing.T) {

	sm := sitemap.New()

	if sm.IsIndex() != false {
		t.Fatalf("New Sitemap should not be Index\n")
	}

	sm.SetIndex(true)

	if sm.IsIndex() != true {
		t.Fatalf("Sitemap should be Index\n")
	}

	sm.SetIndex(false)

	if sm.IsIndex() != false {
		t.Fatalf("Sitemap should not be Index\n")
	}

}

func TestSitemapGetURL(t *testing.T) {

	sm := sitemap.New()
	sm.AppendURL(sitemap.URL{Location: "https://example.com"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/two", LastModification: "2024-01-02"})

	u := sm.GetURL("http://example.com")
	if u != nil {
		t.Fatalf("Found an invalid URL: %#v\n", *u)
	}

	u = sm.GetURL("https://example.com")
	if u == nil {
		t.Fatalf("Not found a valid URL\n")
	}
}

func TestSitemapSortByLocation(t *testing.T) {

	sm := sitemap.New()
	sm.AppendURL(sitemap.URL{Location: "https://example.com/b"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/c", LastModification: "2024-01-01"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/a", LastModification: "2024-01-02"})

	sm.SortByLocation()

	if sm.URL[0].Location != "https://example.com/a" {
		t.Fatalf("The first elem should be /a, got: %s\n", sm.URL[0].Location)
	}

	if sm.URL[1].Location != "https://example.com/b" {
		t.Fatalf("The second elem should be /b, got: %s\n", sm.URL[1].Location)
	}

	if sm.URL[2].Location != "https://example.com/c" {
		t.Fatalf("The third elem should be /c, got: %s\n", sm.URL[2].Location)
	}

}

func TestRemoveURL(t *testing.T) {

	sm := sitemap.New()
	sm.AppendURL(sitemap.URL{Location: "https://example.com/0"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/1"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/2"})
	sm.AppendURL(sitemap.URL{Location: "https://example.com/3"})

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

	sm.RemoveURL("https://example.com/1")
	u1 := sm.GetURL("https://example.com/1")
	if u1 != nil {
		t.Fatalf("Failed to remove the /1 url\n")
	}

}
