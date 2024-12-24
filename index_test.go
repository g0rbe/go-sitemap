package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestIndex(t *testing.T) {

	i1 := sitemap.NewIndex([]*sitemap.Entry{
		sitemap.NewEntry(sitemap.NewLoc("https://example.com/"), sitemap.NewLastMod("2024-01-01")),
		sitemap.NewEntry(sitemap.NewLoc("https://example.com/one"), sitemap.NewLastMod("2024-01-02")),
		sitemap.NewEntry(sitemap.NewLoc("https://example.com/two"), sitemap.NewLastMod("2024-01-03")),
	})

	data, err := i1.ToXML()
	if err != nil {
		t.Fatalf("Failed to marshal to XML: %s\n", err)
	}

	t.Logf("\n%s\n", data)

	i2 := new(sitemap.Index)

	err = xml.Unmarshal(data, i2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if len(i1.Entries) != len(i2.Entries) {
		t.Fatalf("Invalid result:\n%#v\n%#v\n", *i1, *i2)
	}

}
