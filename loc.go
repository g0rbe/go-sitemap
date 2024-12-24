package sitemap

import "encoding/xml"

// Loc is the URL of the page. This URL must begin with the protocol (such as http) and end with a trailing slash, if your web server requires it. This value must be less than 2,048 characters.
//
//	<loc>http://www.example.com/</loc>
type Loc struct {
	XMLName xml.Name `xml:"loc"`
	Value   []byte   `xml:",chardata"`
}

// NewLoc returns a new Loc with the given value v.
//
// This function sets the XMLName to "loc".
func NewLoc(v string) *Loc {
	return &Loc{XMLName: xml.Name{Local: "loc"}, Value: []byte(v)}
}

func (l *Loc) String() string {
	return string(l.Value)
}
