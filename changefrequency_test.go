package sitemap_test

import (
	"encoding/xml"
	"testing"

	"github.com/g0rbe/go-sitemap"
)

func TestUnmarshalChangeFrequency(t *testing.T) {

	var changefreq sitemap.ChangeFrequency = sitemap.ChangeFreqAlways

	data, err := xml.Marshal(changefreq)
	if err != nil {
		t.Fatalf("Filed to marshal: %s\n", err)
	}

	t.Logf("Data: %s\n", data)

	var r sitemap.ChangeFrequency

	err = xml.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Filed to unmarshal: %s\n", err)
	}

	if r != changefreq {
		t.Fatalf("FAIL: invalid data: want %s, got %s\n", changefreq, r)
	}
}
