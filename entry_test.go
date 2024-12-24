package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestEntry(t *testing.T) {

	e1 := sitemap.NewEntry(sitemap.NewLoc("https://example.com/"), nil)

	data, err := xml.MarshalIndent(e1, "", "    ")
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", data)

	e2 := new(sitemap.Entry)

	err = xml.Unmarshal(data, e2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if e1.String() != e2.String() {
		t.Fatalf("Invalid result: %s / %s\n", e1, e2)
	}

}
