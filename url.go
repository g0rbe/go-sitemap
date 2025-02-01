package sitemap

import (
	"net/url"
	"strconv"
	"time"
)

// URL is the parent tag for each URL entry. The remaining tags are children of this tag.
//
// Example:
//
//	<url>
//	  <loc>http://www.example.com/</loc>
//	  <lastmod>2005-01-01</lastmod>
//	  <changefreq>monthly</changefreq>
//	  <priority>0.8</priority>
//	</url>
type URL struct {
	Location         string `xml:"loc"`
	LastModification string `xml:"lastmod,omitempty"`
	ChangeFrequency  string `xml:"changefreq,omitempty" `
	Priority         string `xml:"priority,omitempty" `
}

func (u URL) String() string {

	return u.Location
}

func (u *URL) GetLocation() (*url.URL, error) {
	return url.Parse(u.Location)
}

func (u *URL) GetLastModification() (time.Time, error) {

	var layout string

	switch len(u.LastModification) {
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

	return time.Parse(layout, u.LastModification)
}

func (u *URL) GetPriority() (float64, error) {
	return strconv.ParseFloat(u.Priority, 64)
}
