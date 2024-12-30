package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestChangeFrequency(t *testing.T) {

	l1 := sitemap.ParseChangeFrequency("daily")

	out, err := xml.Marshal(l1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<changefreq>daily</changefreq>")) {
		t.Fatalf("Invalid string: %s\n", out)
	}

	var l2 = new(sitemap.ChangeFrequency)

	err = xml.Unmarshal(out, l2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if l1.String() != l2.String() {
		t.Fatalf("Invalid result: %s / %s\n", l1, l2)
	}
}
