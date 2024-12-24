package sitemap_test

import (
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURL(t *testing.T) {

	u1 := sitemap.NewURL(sitemap.NewLoc("https://example.com"), nil, nil, nil)

	out, err := xml.MarshalIndent(u1, "", "    ")
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", out)

	var u2 = new(sitemap.URL)

	err = xml.Unmarshal(out, u2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if u1.String() != u2.String() {
		t.Fatalf("Invalid result: %s / %s\n", u1, u2)
	}

}
