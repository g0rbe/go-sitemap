package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestLastMod(t *testing.T) {

	l1 := sitemap.NewLastMod("2024-01-01")

	out, err := xml.Marshal(l1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("%s\n", out)

	var l2 = new(sitemap.LastMod)

	err = xml.Unmarshal(out, l2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if l1.String() != l2.String() {
		t.Fatalf("Invalid result: %s / %s\n", l1, l2)
	}
}
