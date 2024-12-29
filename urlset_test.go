package sitemap_test

import (
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

	t.Logf("\n%s\n", data)

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

	t.Logf("Length: %d\n", len(u.URLs))
}

func ExampleFetchURLSet() {

	u, err := sitemap.FetchURLSet("https://gorbe.io/en/sitemap.xml")
	if err != nil {
		// Handle error
	}

	fmt.Printf("Length is %d\n", len(u.URLs))
}
