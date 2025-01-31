package sitemap

import (
	"encoding/xml"
)

// Index encapsulates information about all of the Sitemaps in the file.
//
//	<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//	  <sitemap>
//	    <loc>http://www.example.com/sitemap1.xml.gz</loc>
//	    <lastmod>2004-10-01T18:23:17+00:00</lastmod>
//	  </sitemap>
//	  <sitemap>
//	    <loc>http://www.example.com/sitemap2.xml.gz</loc>
//	    <lastmod>2005-01-01</lastmod>
//	  </sitemap>
//	</sitemapindex>
type Index struct {
	XMLName   xml.Name `xml:"sitemapindex"`
	NameSpace string   `xml:"xmlns,attr"`
	Sitemap   []URL    `xml:"sitemap"`
}

// NewINdex returns a new Index with the given Sitemap Entries.
//
// This function sets the XMLName to "sitemapindex".
func NewIndex(sitemaps ...URL) *Index {
	return &Index{NameSpace: XMLNameSpace, Sitemap: sitemaps}
}
