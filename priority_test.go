package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestPriority(t *testing.T) {

	p1 := sitemap.NewPriority("0.5")

	out, err := xml.Marshal(p1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<priority>0.5</priority>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}

	var p2 = new(sitemap.Priority)

	err = xml.Unmarshal(out, p2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if p1.String() != p2.String() {
		t.Fatalf("Invalid result: %s / %s\n", p1, p2)
	}
}
