package sitemap_test

import (
	"encoding/xml"
	"testing"

	"github.com/g0rbe/go-sitemap"
)

func TestUnmasrhalPriority(t *testing.T) {

	var priority sitemap.Priority = 0.7

	data, err := xml.Marshal(priority)
	if err != nil {
		t.Fatalf("Filed to marshal: %s\n", err)
	}

	t.Logf("Data: %s\n", data)

	var result sitemap.Priority

	err = xml.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if result != priority {
		t.Fatalf("FAIL: invalid result:got %s, want %s\n", result, priority)
	}
}
