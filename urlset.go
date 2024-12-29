package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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
	URLs []*URL `xml:"url"`
}

// NewURLSet returns a new URLSet with the given URLs.
//
// This function sets the XMLName to "urlset" and XMLNS to "http://www.sitemaps.org/schemas/sitemap/0.9".
func NewURLSet(urls []*URL) *URLSet {
	return &URLSet{URLs: urls}
}

// ReadURLSet reads the Sitemap from r.
//
// If r contains Sitemap INdex, returns an empty URLSet.
func ReadURLSet(r io.Reader) (*URLSet, error) {

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	u := new(URLSet)

	err = xml.Unmarshal(data, u)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return u, nil
}

// FetchURLSet fetches the Sitemap from url.
//
// If r contains Sitemap INdex, returns an empty URLSet.
func FetchURLSet(url string) (*URLSet, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ReadURLSet(resp.Body)
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

func (u *URLSet) ToTXT() ([]byte, error) {

	buf := new(bytes.Buffer)

	_, err := buf.Write([]byte(xml.Header))
	if err != nil {
		return nil, fmt.Errorf("failed to write XML header: %w", err)
	}

	for i := range u.URLs {
		_, err = buf.WriteString(u.URLs[i].Loc.String())
		if err != nil {
			return nil, fmt.Errorf("failed to write data: %w", err)
		}

		err = buf.WriteByte('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to write newline: %w", err)
		}
	}

	return buf.Bytes(), nil
}

func (u *URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "urlset"
	if start.Name.Local != "urlset" {
		start.Name.Local = "urlset"
	}

	// Append "xmlns" attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: XMLNameSpace})

	v := struct {
		URLs []*URL `xml:"url"`
	}{
		URLs: u.URLs,
	}

	return e.EncodeElement(v, start)
}
