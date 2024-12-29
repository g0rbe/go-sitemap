package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURL(t *testing.T) {

	u1 := sitemap.NewURL(sitemap.NewLocation("https://example.com"), sitemap.NewLastMod("1970-01-01"), nil, nil)

	out, err := xml.Marshal(u1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<url><loc>https://example.com</loc><lastmod>1970-01-01</lastmod></url>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}

	var u2 = new(sitemap.URL)

	err = xml.Unmarshal(out, u2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if u1.String() != u2.String() {
		t.Fatalf("Invalid result: %s / %s\n", u1, u2)
	}

}
