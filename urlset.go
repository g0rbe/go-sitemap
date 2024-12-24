package sitemap

import "encoding/xml"

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
	XMLName xml.Name `xml:"urlset"`
	NS      []byte   `xml:"xmlns,attr"`
	URLs    []*URL   `xml:"url"`
}

// NewURLSet returns a new URLSet with the given URLs.
//
// This function sets the XMLName to "urlset".
func NewURLSet(urls []*URL) *URLSet {
	return &URLSet{XMLName: xml.Name{Local: "urlset"}, NS: []byte("http://www.sitemaps.org/schemas/sitemap/0.9"), URLs: urls}
}
