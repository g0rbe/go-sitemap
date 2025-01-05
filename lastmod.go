package sitemap

import (
	"encoding/xml"
	"time"
)

// The date of last modification of the page. This date should be in W3C Datetime format. This format allows you to omit the time portion, if desired, and use YYYY-MM-DD.
//
// Note that the date must be set to the date the linked page was last modified, not when the sitemap is generated.
//
// Note also that this tag is separate from the If-Modified-Since (304) header the server can return, and search engines may use the information from both sources differently.
//
// Example:
//
//	<lastmod>2005-01-01</lastmod>
type LastModification string

// NewLastMod returns a new LastMod with the given value v.
// If v is an empty string("") or zero Time, returns nil.
func NewLastModification[T string | time.Time](v T) *LastModification {

	switch t := any(v).(type) {
	case string:
		if t == "" {
			return nil
		}
		return (*LastModification)(&t)

	case time.Time:
		if t.IsZero() {
			return nil
		}
		f := t.Format(time.RFC3339)
		return (*LastModification)(&f)
	default:
		return nil
	}

}

func (l *LastModification) String() string {
	return string(*l)
}

func (l *LastModification) Time() (time.Time, error) {

	var layout string

	switch len(*l) {
	case 4:
		layout = "2006"
	case 7:
		layout = "2006-01"
	case 10:
		layout = "2006-01-02"
	case 20:
		layout = "2006-01-02T15:04:05Z"
	case 24:
		layout = "2006-01-02T15:04:05.999Z"
	case 25:
		layout = "2006-01-02T15:04:05.999999999Z07:00"
	default:
		layout = time.RFC3339Nano
	}

	return time.Parse(layout, l.String())

}
func (l *LastModification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// CHange the start and end tag to "lastmod"
	if start.Name.Local != "lastmod" {
		start.Name.Local = "lastmod"
	}

	return e.EncodeElement(l.String(), start)
}
