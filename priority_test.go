package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestPriorityXMLMarhsal(t *testing.T) {

	prio := sitemap.Priority("0.5")

	out, err := xml.Marshal(prio)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<Priority>0.5</Priority>")) {
		t.Fatalf("Invalid data: %s\n", out)
	}
}

func TestPriorityXMLUnmarshal(t *testing.T) {

	prio := new(sitemap.Priority)

	err := xml.Unmarshal([]byte("<Priority>0.5</Priority>"), prio)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if prio.String() != "0.5" {
		t.Fatalf("Invalid result: %s\n", prio)
	}
}
