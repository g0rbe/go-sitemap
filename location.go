package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/url"
)

// Location type used for the <loc> field
type Location []byte

func (l Location) String() string {
	return string(l)
}

func (l Location) Bytes() []byte {
	return l
}

func (l Location) IsEmpty() bool {
	return len(l) == 0
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (l *Location) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	u := new([]byte)

	err := d.DecodeElement(u, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Location: %w", err)
	}

	*u = bytes.TrimSpace(*u)

	_, err = url.Parse(string(*u))
	if err != nil {
		return fmt.Errorf("invalid Location: \"%s\" (%w)", *u, err)
	}

	*l = *(*Location)(u)

	return nil
}

func (l Location) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "loc"

	return e.EncodeElement(l.Bytes(), start)
}

func (l Location) EqualTo(j Location) bool {
	return bytes.Equal(l, j)
}
