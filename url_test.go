package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/g0rbe/go-sitemap"
)

func TestUnmarshalURL(t *testing.T) {

	var url sitemap.URL = sitemap.URL{
		Loc:        sitemap.Location("https://gorbe.io/about"),
		LastMod:    sitemap.LastModification(time.Now().Unix()),
		ChangeFreq: sitemap.ChangeFreqAlways,
		Priority:   sitemap.Priority(1.0),
	}

	data, err := xml.Marshal(url)
	if err != nil {
		t.Fatalf("Filed to marshal: %s\n", err)
	}

	t.Logf("Data: %s\n", data)

	var result sitemap.URL

	err = xml.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if !bytes.Equal(url.Loc, result.Loc) {
		t.Fatalf("FAIL: invalid URL result: got %s, want %s\n", result.Loc, url.Loc)
	}
}
