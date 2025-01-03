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
	Location   *Location         `xml:"loc"`
	LastMod    *LastModification `xml:"lastmod,omitempty"`
	ChangeFreq *ChangeFrequency  `xml:"changefreq,omitempty"`
	Priority   *Priority         `xml:"priority,omitempty"`
	Comment    []byte            `xml:",comment"`
}

// NewURL returns a new URL with the given fields set.
// If any field is nil, it will be omotted.
//
// This function sets the XMLName to "url".
func NewURL(loc *Location) *URL {
	return &URL{Location: loc}
}

// SetLastmodification lastmod in URL u and returns u.
//
// If URL u is nil, returns nil.
func (u *URL) SetLastmodification(lastmod *LastModification) *URL {

	if u == nil {
		return nil
	}

	u.LastMod = lastmod
	return u
}

func (u *URL) SetChangeFrequency(changefreq *ChangeFrequency) *URL {

	if u == nil {
		return nil
	}

	u.ChangeFreq = changefreq
	return u
}

func (u *URL) SetPriority(priority *Priority) *URL {

	if u == nil {
		return nil
	}

	u.Priority = priority
	return u
}

func (u *URL) String() string {
	buf := new(strings.Builder)

	buf.WriteString(u.Location.String())

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

func (u *URL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "url" if not set to "url" for URLSet or "sitemap" for Index.
	if start.Name.Local != "url" && start.Name.Local != "sitemap" {
		start.Name.Local = "url"
	}

	v := struct {
		Loc        *Location         `xml:"loc"`
		LastMod    *LastModification `xml:"lastmod,omitempty"`
		ChangeFreq *ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority   *Priority         `xml:"priority,omitempty"`
		Comment    []byte            `xml:",comment"`
	}{
		Loc:        u.Location,
		LastMod:    u.LastMod,
		ChangeFreq: u.ChangeFreq,
		Priority:   u.Priority,
		Comment:    u.Comment,
	}

	return e.EncodeElement(v, start)
}
