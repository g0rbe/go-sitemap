package sitemap_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/g0rbe/go-sitemap"
)

func TestUnmasrhalLastModification(t *testing.T) {

	v := sitemap.LastModification(time.Now().Unix())

	data, err := xml.Marshal(v)
	if err != nil {
		t.Fatalf("Filed to marshal: %s\n", err)
	}

	t.Logf("Data: %s\n", data)

	var r sitemap.LastModification

	err = xml.Unmarshal(data, &r)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if r.String() != v.String() {
		t.Fatalf("FAIL: invalid data: want %s, got %s\n", v, r)
	}
}
