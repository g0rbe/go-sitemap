package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// LastModification is the <lastmod> field stored in Unix timestamp.
type LastModification int64

func ParseLastModification(s string) (LastModification, error) {

	var layout string

	switch len(s) {
	case 10:
		layout = time.DateOnly
	case 20:
		layout = "2006-01-02T15:04:05Z"
	case 24:
		layout = "2006-01-02T15:04:05.999Z"
	case 25:
		layout = "2006-01-02T15:04:05.999999999Z07:00"
	default:
		layout = time.RFC3339Nano
	}

	v, err := time.Parse(layout, s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s: %w", s, err)
	}

	return LastModification(v.Unix()), nil

}
func (t LastModification) String() string {
	return time.Unix(int64(t), 0).Format(time.RFC3339Nano)
}

func (t LastModification) IsEmpty() bool {
	return t == 0
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (t *LastModification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string

	err := d.DecodeElement(&s, &start)
	if err != nil {
		return fmt.Errorf("failed to decode LastModification: %w", err)
	}

	v, err := ParseLastModification(strings.TrimSpace(s))

	*t = v

	return err
}

func (l LastModification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "lastmod"

	return e.EncodeElement(l.String(), start)
}
