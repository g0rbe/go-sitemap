package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
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
	XMLName xml.Name `xml:"sitemapindex"`
	NS      []byte   `xml:"xmlns,attr"`
	Entries []*Entry `xml:"sitemap"`
}

// NewINdex returns a new Index with the given Sitemap Entries.
//
// This function sets the XMLName to "sitemapindex".
func NewIndex(entries []*Entry) *Index {
	return &Index{
		XMLName: xml.Name{Local: "sitemapindex"},
		NS:      []byte("http://www.sitemaps.org/schemas/sitemap/0.9"),
		Entries: entries}
}

func (i *Index) ToXML() ([]byte, error) {

	data, err := xml.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}

	buf := new(bytes.Buffer)

	_, err = buf.Write([]byte(xml.Header))
	if err != nil {
		return nil, fmt.Errorf("failed to write XML header: %w", err)
	}

	_, err = buf.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data: %w", err)
	}

	return buf.Bytes(), nil
}
