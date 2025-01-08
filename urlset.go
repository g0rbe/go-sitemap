package sitemap

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
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
	URLs []URL `xml:"url" json:"url"`
}

// NewURLSet returns a new URLSet with the given URLs.
func NewURLSet(urls ...URL) *URLSet {

	s := new(URLSet)

	if len(urls) > 0 {
		s.URLs = append(s.URLs, urls...)
	}

	return s
}

// ReadURLSet reads the Sitemap from r.
//
// If r contains Sitemap Index, returns ErrSitemapIndex.
func ParseURLSet(data []byte) (*URLSet, error) {

	if IsIndex(data) {
		return nil, ErrSitemapIndex
	}

	u := new(URLSet)

	err := xml.Unmarshal(data, u)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return u, nil
}

func (u *URLSet) ToXML() ([]byte, error) {

	data, err := xml.Marshal(u)
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

func (u *URLSet) ToXMLIndent() ([]byte, error) {

	data, err := xml.MarshalIndent(u, "", "\t")
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

func (u *URLSet) ToTXT() ([]byte, error) {

	buf := new(bytes.Buffer)

	for i := range u.URLs {

		// Write Location + "\n"
		_, err := buf.WriteString(u.URLs[i].Location.String() + "\n")
		if err != nil {
			return nil, fmt.Errorf("failed to write %s: %w", *u.URLs[i].Location, err)
		}
	}

	return buf.Bytes(), nil
}

func (u *URLSet) ToJSON() ([]byte, error) {

	v := struct {
		URLSet *URLSet `json:"urlset"`
	}{
		URLSet: u,
	}

	return json.Marshal(v)

}

func (u *URLSet) ToJSONIndent() ([]byte, error) {

	v := struct {
		URLSet *URLSet `json:"urlset"`
	}{
		URLSet: u,
	}

	return json.MarshalIndent(v, "", "\t")
}

func (u *URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "urlset"
	if start.Name.Local != "urlset" {
		start.Name.Local = "urlset"
	}

	// Append "xmlns" attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: XMLNameSpace})

	v := struct {
		URLs []URL `xml:"url"`
	}{
		URLs: u.URLs,
	}

	return e.EncodeElement(v, start)
}
