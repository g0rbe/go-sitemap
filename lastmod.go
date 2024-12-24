package sitemap

import "encoding/xml"

// The date of last modification of the page. This date should be in W3C Datetime format. This format allows you to omit the time portion, if desired, and use YYYY-MM-DD.
//
// Note that the date must be set to the date the linked page was last modified, not when the sitemap is generated.
//
// Note also that this tag is separate from the If-Modified-Since (304) header the server can return, and search engines may use the information from both sources differently.
//
//	<lastmod>2005-01-01</lastmod>
type LastMod struct {
	XMLName xml.Name `xml:"lastmod"`
	Value   []byte   `xml:",chardata"`
}

// NewLastMod returns a new LastMod with the given value v.
// If v is an emty string(""), returns nil.
//
// This function sets the XMLName to "lastmod".
func NewLastMod(v string) *LastMod {
	if v == "" {
		return nil
	}
	return &LastMod{XMLName: xml.Name{Local: "lastmod"}, Value: []byte(v)}
}

func (l *LastMod) String() string {
	return string(l.Value)
}
