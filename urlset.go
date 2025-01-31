package sitemap

import (
	"encoding/xml"
	"errors"
)

var (
	ErrSitemapIndex = errors.New("sitemap is index")
)

// URLSet encapsulates the file and references the current protocol standard.
//
//	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//	  <url>
//	    <loc>http://www.example.com/</loc>
//	    <lastmod>2005-01-01</lastmod>
//	    <changefreq>monthly</changefreq>
//	    <priority>0.8</priority>
//	  </url>
//	</urlset>
type URLSet struct {
	XMLName   xml.Name `xml:"urlset"`
	NameSpace string   `xml:"xmlns,attr"`
	URL       []URL    `xml:"url"`
}

// NewURLSet returns a new URLSet with the given URLs.
func NewURLSet(urls ...URL) *URLSet {
	return &URLSet{NameSpace: XMLNameSpace, URL: urls}
}
