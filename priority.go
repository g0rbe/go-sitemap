package sitemap

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// Priority type used for the priority field
type Priority float64

func (p *Priority) String() string {
	return strconv.FormatFloat(float64(*p), 'g', -1, 64)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (p *Priority) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	s := new(string)

	err := d.DecodeElement(s, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Priority: %w", err)
	}

	f, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Priority %s: %w", *s, err)

	}

	if f < 0.0 || f > 1.0 {
		return fmt.Errorf("invalid value for priority: %f", f)
	}

	*p = *(*Priority)(&f)

	return nil
}
