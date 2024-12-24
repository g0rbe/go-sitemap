package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURLSet(t *testing.T) {

	u1 := sitemap.NewURLSet(
		[]*sitemap.URL{
			sitemap.NewURL(sitemap.NewLoc("https://example.com"), nil, nil, nil),
			sitemap.NewURL(sitemap.NewLoc("https://example.com/two"), sitemap.NewLastMod("2024-01-02"), nil, nil),
		})

	data, err := xml.MarshalIndent(u1, "", "    ")
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
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
