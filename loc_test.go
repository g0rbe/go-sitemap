package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestLoc(t *testing.T) {

	l1 := sitemap.NewLoc("https://gorbe.io/test")

	out, err := xml.Marshal(l1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("%s\n", out)

	var l2 = new(sitemap.Loc)

	err = xml.Unmarshal(out, l2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if l1.String() != l2.String() {
		t.Fatalf("Invalid result: %s / %s\n", l1, l2)
	}
}
