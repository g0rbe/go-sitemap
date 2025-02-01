package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestURLXMLMarshal(t *testing.T) {

	u := sitemap.URL{Location: "https://example.com", LastModification: "1970-01-01"}

	out, err := xml.Marshal(u)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<URL><loc>https://example.com</loc><lastmod>1970-01-01</lastmod></URL>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}
}

func TestURLXMLUnmarshal(t *testing.T) {

	var u = new(sitemap.URL)

	err := xml.Unmarshal([]byte("<URL><loc>https://example.com</loc><lastmod>1970-01-01</lastmod></URL>"), u)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if u.String() != "https://example.com" {
		t.Fatalf("Invalid result: %s\n", u)
	}
}

func ExampleURL() {

	u := sitemap.URL{Location: "https://example.com", LastModification: "1970-01-01", ChangeFrequency: "never", Priority: "1.0"}

	fmt.Printf("%s\n", u)
	// Output: https://example.com
}
