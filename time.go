package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type Time time.Time

func (t *Time) String() string {
	return time.Time(*t).Format(time.RFC3339Nano)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	s := new(string)

	err := d.DecodeElement(s, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Time: %w", err)
	}

	*s = strings.TrimSpace(*s)
	var layout string

	switch len(*s) {
	case 10:
		layout = time.DateOnly
	case 20:
		layout = "2006-01-02T15:04:05Z"
	case 24:
		layout = "2006-01-02T15:04:05.999Z"
	case 25:
		layout = "2006-01-02T15:04:05.999999999Z07:00"
	default:
		return fmt.Errorf("unknown Time layout: %s", *s)
	}

	v, err := time.Parse(layout, *s)
	if err != nil {
		return fmt.Errorf("failed to parse Time %s: %w", *s, err)
	}

	*t = *(*Time)(&v)

	return nil
}
