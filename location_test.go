package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestLocationXMLMarhsal(t *testing.T) {

	loc := sitemap.Location("https://gorbe.io/test")

	out, err := xml.Marshal(loc)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<Location>https://gorbe.io/test</Location>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}
}

func TestLocationXMLUnmarhsal(t *testing.T) {

	loc := new(sitemap.Location)

	err := xml.Unmarshal([]byte("<Location>https://gorbe.io/test</Location>"), loc)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if loc.String() != "https://gorbe.io/test" {
		t.Fatalf("Invalid result: %s\n", loc)
	}
}
