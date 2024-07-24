package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/g0rbe/go-sitemap"
)

func TestUnmarshalLocation(t *testing.T) {

	loc := sitemap.Location("https://gorbe.io/about")

	data, err := xml.Marshal(loc)
	if err != nil {
		t.Fatalf("Filed to marshal: %s\n", err)
	}

	t.Logf("Data: %s\n", data)

	var result sitemap.Location

	err = xml.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if !bytes.Equal(result, loc) {
		t.Fatalf("FAIL: invalid result:got %s, want %s\n", result, loc)
	}
}
