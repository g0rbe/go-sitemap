package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestLastModificationXMLMarshal(t *testing.T) {

	lastmod := sitemap.LastModification("2024-01-01")

	out, err := xml.Marshal(lastmod)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<LastModification>2024-01-01</LastModification>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}
}

func TestLastModificationXMLUnmarshal(t *testing.T) {

	lastmod := new(sitemap.LastModification)

	err := xml.Unmarshal([]byte("<LastModification>2024-01-01</LastModification>"), lastmod)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if lastmod.String() != "2024-01-01" {
		t.Fatalf("Invalid result: %s\n", lastmod)
	}
}
