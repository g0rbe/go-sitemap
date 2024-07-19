package sitemap_test

import (
	"fmt"
	"testing"

	"github.com/g0rbe/go-sitemap"
)

var (
	TestURLSet = "https://gorbe.io/sitemap.xml"
	TestIndex  = "https://drszokek.hu/sitemap_index.xml"
)

var TestSites = []string{
	"https://gorbe.io/sitemap.xml", "https://thehackernews.com/sitemap.xml",
	"https://www.hubspot.com/sitemap.xml", "https://www.apple.com/sitemap.xml",
	"https://www.atlassian.com/sitemap.xml",
}

func TestFetchFormat(t *testing.T) {

	for i := range TestSites {

		f, _, err := sitemap.FetchFormat(TestSites[i])
		if err != nil {
			t.Errorf("%s -> %s\n", TestSites[i], err)
			continue
		}

		t.Logf("%s -> %s\n", TestSites[i], f)
	}
}

func TestFetch(t *testing.T) {

	for i := range TestSites {

		v, err := sitemap.Fetch(TestSites[i])
		if err != nil || len(v) == 0 {
			t.Errorf("%s -> %s\n", TestSites[i], err)
			continue
		}

		t.Logf("%s -> %d\n", TestSites[i], len(v))

	}
}

func ExampleFetchFormat() {

	f, _, err := sitemap.FetchFormat("https://gorbe.io/sitemap.xml")
	if err != nil {
		// handle error
	}

	fmt.Printf("%s\n", f) // "urlset"
}

func ExampleGetFormat() {

	data := []byte(`
<?xml version="1.0" encoding="UTF-8"?>

<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">

   <url>

      <loc>http://www.example.com/</loc>

      <lastmod>2005-01-01</lastmod>

      <changefreq>monthly</changefreq>

      <priority>0.8</priority>

   </url>

</urlset>`)

	f, err := sitemap.GetFormat(data)
	if err != nil {
		// handle error
	}

	fmt.Printf("%s\n", f) // "urlset"

}

func ExampleFetch() {

	sm, err := sitemap.Fetch("https://gorbe.io/sitemap.xml")
	if err != nil {
		// handle error
	}

	for i := range sm {
		fmt.Printf("%s\n", sm[i].String())
	}
}
