package sitemap

import (
	"encoding/xml"
	"strings"
)

// URL is the parent tag for each URL entry. The remaining tags are children of this tag.
//
//	<url>
//	  <loc>http://www.example.com/</loc>
//	  <lastmod>2005-01-01</lastmod>
//	  <changefreq>monthly</changefreq>
//	  <priority>0.8</priority>
//	</url>
type URL struct {
	XMLName    xml.Name    `xml:"url"`
	Loc        *Loc        `xml:"loc"`
	LastMod    *LastMod    `xml:"lastmod,omitempty"`
	ChangeFreq *ChangeFreq `xml:"changefreq,omitempty"`
	Priority   *Priority   `xml:"priority,omitempty"`
	Comment    []byte      `xml:",comment"`
}

// NewURL returns a new URL with the given fields set.
// If any field is nil, it will be omotted.
//
// This function sets the XMLName to "url".
func NewURL(loc *Loc, lastmod *LastMod, changefreq *ChangeFreq, prio *Priority) *URL {
	return &URL{
		XMLName:    xml.Name{Local: "url"},
		Loc:        loc,
		LastMod:    lastmod,
		ChangeFreq: changefreq,
		Priority:   prio}
}

func (u *URL) String() string {
	buf := new(strings.Builder)

	buf.WriteString(u.Loc.String())

	if u.LastMod != nil {
		buf.WriteByte(' ')
		buf.WriteString(u.LastMod.String())
	}

	if u.ChangeFreq != nil {
		buf.WriteByte(' ')
		buf.WriteString(u.ChangeFreq.String())
	}

	if u.Priority != nil {
		buf.WriteByte(' ')
		buf.WriteString(u.Priority.String())
	}

	return buf.String()
}
