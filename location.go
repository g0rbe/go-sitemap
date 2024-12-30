package sitemap

import (
	"encoding/xml"
	"net/url"
)

// Location is the URL of the page. This URL must begin with the protocol (such as http) and end with a trailing slash, if your web server requires it. This value must be less than 2,048 characters.
//
// Example:
//
//	<loc>http://www.example.com/</loc>
type Location string

// NewLoc returns a new Loc with the given value v.
//
// This function sets the XMLName to "loc".
func NewLocation[T string | url.URL](v T) *Location {

	switch t := any(v).(type) {
	case string:
		return (*Location)(&t)
	case url.URL:
		r := t.String()
		return (*Location)(&r)
	default:
		return nil
	}
}

func (l *Location) String() string {
	return string(*l)
}

func (l *Location) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := new(string)

	err := d.DecodeElement(v, &start)
	if err != nil {
		return err
	}

	*l = Location(*v)

	return nil
}

func (l *Location) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// CHange the start and end tag to "loc"
	if start.Name.Local != "loc" {
		start.Name.Local = "loc"
	}

	return e.EncodeElement(l.String(), start)
}
