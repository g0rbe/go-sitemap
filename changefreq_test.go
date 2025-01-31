package sitemap_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"git.gorbe.io/go/sitemap"
)

func TestChangeFrequencyXMLMarshal(t *testing.T) {

	freq := sitemap.ChangeFrequencyDaily

	out, err := xml.Marshal(freq)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	if !bytes.Equal(out, []byte("<ChangeFrequency>daily</ChangeFrequency>")) {
		t.Fatalf("Invalid string: %s\n", out)
	}
}

func TestChangeFrequencyXMLUnmarshal(t *testing.T) {

	var freq = new(sitemap.ChangeFrequency)

	err := xml.Unmarshal([]byte("<ChangeFrequency>daily</ChangeFrequency>"), freq)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if freq.String() != "daily" {
		t.Fatalf("Invalid result: %s\n", freq)
	}
}

// func TestChangeFrequencyIsValid(t *testing.T) {

// 	freq := sitemap.ChangeFrequencyAlways

// 	if !freq.IsValid() {
// 		t.Fatalf("%s is marked as invalid\n", freq)
// 	}

// 	freq = sitemap.ChangeFrequency("sometimes")
// 	if freq.IsValid() {
// 		t.Fatalf("%s is marked as valid\n", freq)
// 	}
// }
