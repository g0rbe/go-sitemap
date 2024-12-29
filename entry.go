package sitemap

import (
	"encoding/xml"
	"strings"
)

// Entry is encapsulates information about an individual Sitemap in Sitemap Index.
//
//	<sitemap>
//	  <loc>http://www.example.com/sitemap1.xml.gz</loc>
//	  <lastmod>2004-10-01T18:23:17+00:00</lastmod>
//	</sitemap>
type Entry struct {
	XMLName xml.Name  `xml:"sitemap"`
	Loc     *Location `xml:"loc"`
	LastMod *LastMod  `xml:"lastmod,omitempty"`
	Comment []byte    `xml:",comment"`
}

// NewEntry returns a new Entry with the given fields set.
//
// This function sets the XMLName to "sitemap".
func NewEntry(loc *Location, lastmod *LastMod) *Entry {
	return &Entry{XMLName: xml.Name{Local: "sitemap"}, Loc: loc, LastMod: lastmod}
}

func (e *Entry) String() string {
	buf := new(strings.Builder)

	buf.WriteString(e.Loc.String())

	if e.LastMod != nil {
		buf.WriteByte(' ')
		buf.WriteString(e.LastMod.String())
	}

	return buf.String()
}
