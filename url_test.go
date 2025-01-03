package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURL(t *testing.T) {

	u1 := sitemap.NewURL(sitemap.NewLocation("https://example.com")).SetLastmodification(sitemap.NewLastModification("1970-01-01"))

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

func ExampleNewURL() {

	u := sitemap.NewURL(sitemap.NewLocation("https://example.com")).
		SetLastmodification(sitemap.NewLastModification("1970-01-01")).
		SetChangeFrequency(&sitemap.ChangeFrequencyNever).
		SetPriority(sitemap.NewPriority("1.0"))

	fmt.Printf("%s\n", u)
	// Output: https://example.com 1970-01-01 never 1.0
}
