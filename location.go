package sitemap

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

// Location type used for the loc field
type Location url.URL

func (l *Location) String() string {
	return (*url.URL)(l).String()
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (l *Location) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	u := new(string)

	err := d.DecodeElement(u, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Location: %w", err)
	}

	v, err := url.Parse(*u)
	if err != nil {
		return fmt.Errorf("failed to parse Location %s: %w", *u, err)
	}

	*l = *(*Location)(v)

	return nil
}
