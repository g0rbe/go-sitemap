package sitemap

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
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
	URLs []*URL `xml:"url"`
	m    *sync.RWMutex
}

// EmptyURLSet returns a new  URLSet without URLs.
func EmptyURLSet() *URLSet {

	s := new(URLSet)
	s.m = new(sync.RWMutex)

	return s
}

// NewURLSet returns a new URLSet with the given URLs.
func NewURLSet(urls ...*URL) *URLSet {

	s := EmptyURLSet()

	if len(urls) > 0 {
		s.URLs = append(s.URLs, urls...)
	}

	return s
}

// ReadURLSet reads the Sitemap from r.
//
// If r contains Sitemap Index, returns ErrSitemapIndex.
func ReadURLSet(r io.Reader) (*URLSet, error) {

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	if SitemapIsIndex(data) {
		return nil, ErrSitemapIndex
	}

	u := EmptyURLSet()

	err = xml.Unmarshal(data, u)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return u, nil
}

// FetchURLSet fetches the Sitemap from url.
//
// If r contains Sitemap Index, returns ErrSitemapIndex.
func FetchURLSet(url string) (*URLSet, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ReadURLSet(resp.Body)
}

// Size returns the number of URL in u.URLs
func (u *URLSet) Size() int {
	u.m.RLock()
	defer u.m.RUnlock()
	return len(u.URLs)
}

func (u *URLSet) GetURL(loc string) *URL {

	u.m.RLock()
	defer u.m.RUnlock()

	for i := range u.URLs {
		if *u.URLs[i].Location == Location(loc) {
			return u.URLs[i]
		}
	}

	return nil
}

func (u *URLSet) SetURL(url *URL) {

	u.m.Lock()
	defer u.m.Unlock()

	for i := range u.URLs {
		if *u.URLs[i].Location == *url.Location {
			u.URLs[i] = url
			return
		}
	}

	u.URLs = append(u.URLs, url)
}

func (u *URLSet) ToXML() ([]byte, error) {

	u.m.RLock()
	defer u.m.RUnlock()

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

	u.m.RLock()
	defer u.m.RUnlock()

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

func (u *URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	u.m.RLock()
	defer u.m.RUnlock()

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
