package sitemap

import (
	"bytes"
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
	Location   Location         `xml:"loc"`
	LastMod    LastModification `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFrequency  `xml:"changefreq,omitempty" `
	Priority   Priority         `xml:"priority,omitempty" `
}

// NewURL returns a new URL with the given fields set.
// If any field is nil, it will be omitted.
func NewURL(loc string) *URL {
	return &URL{Location: bytes.Clone(Location(loc))}
}

// SetLastmodification clones lastmod to u.LastMod and returns u.
//
// If URL u is nil, returns nil.
func (u *URL) SetLastmodification(lastmod string) *URL {

	if u == nil {
		return nil
	}

	u.LastMod = bytes.Clone(LastModification(lastmod))
	return u
}

// SetChangeFrequency clones changefreq to u.ChangeFreq and returns u.
//
// If URL u is nil, returns nil.
func (u *URL) SetChangeFrequency(changefreq string) *URL {

	if u == nil {
		return nil
	}

	u.ChangeFreq = bytes.Clone(ChangeFrequency(changefreq))
	return u
}

// SetPriority clones priority to u.Priority and returns u.
//
// If URL u is nil, returns nil.
func (u *URL) SetPriority(priority string) *URL {

	if u == nil {
		return nil
	}

	u.Priority = bytes.Clone(Priority(priority))
	return u
}

func (u URL) String() string {

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
