package sitemap

import (
	"bytes"
	"encoding/json"
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
	Sitemaps []URL `xml:"sitemap" json:"sitemap"`
}

// NewINdex returns a new Index with the given Sitemap Entries.
//
// This function sets the XMLName to "sitemapindex".
func NewIndex(sitemaps ...URL) *Index {
	return &Index{Sitemaps: sitemaps}
}

// ReadURLSet reads the Sitemap from r.
//
// If r contains Sitemap Index, returns ErrSitemapIndex.
func ParseIndex(data []byte) (*Index, error) {

	if !IsIndex(data) {
		return nil, fmt.Errorf("not index")
	}

	i := new(Index)

	err := xml.Unmarshal(data, i)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return i, nil
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

func (i *Index) ToJSON() ([]byte, error) {

	v := struct {
		Index *Index `json:"sitemapindex"`
	}{
		Index: i,
	}

	return json.Marshal(v)
}

func (i *Index) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "sitemapindex"
	if start.Name.Local != "sitemapindex" {
		start.Name.Local = "sitemapindex"
	}

	// Append "xmlns" attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: XMLNameSpace})

	v := struct {
		Sitemaps []URL `xml:"sitemap"`
	}{
		Sitemaps: i.Sitemaps,
	}

	return e.EncodeElement(v, start)
}
