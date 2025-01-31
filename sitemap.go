package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

const XMLNameSpace = "http://www.sitemaps.org/schemas/sitemap/0.9"

func IsIndex(data []byte) bool {
	return bytes.Contains(data, []byte("sitemapindex"))
}

type Sitemap struct {
	URL     []URL
	isIndex bool
	m       *sync.RWMutex
}

func New() *Sitemap {
	s := new(Sitemap)
	s.m = new(sync.RWMutex)
	return s
}

func ParseXMLIndex(data []byte) (*Sitemap, error) {
	sm := New()
	sm.isIndex = true

	index := NewIndex()

	err := xml.Unmarshal(data, index)
	if err != nil {
		return nil, err
	}

	sm.URL = index.Sitemap

	return sm, nil
}

func ParseXMLURLSet(data []byte) (*Sitemap, error) {

	sm := New()
	sm.isIndex = true

	urlset := NewURLSet()

	err := xml.Unmarshal(data, urlset)
	if err != nil {
		return nil, err
	}

	sm.URL = urlset.URL

	return sm, nil
}

// Parse reads the Sitemap from data.
func ParseXML(data []byte) (*Sitemap, error) {

	if IsIndex(data) {
		return ParseXMLIndex(data)
	} else {
		return ParseXMLURLSet(data)
	}
}

// Read reads the Sitemap from r.
func ReadXML(r io.Reader) (*Sitemap, error) {

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	return ParseXML(data)
}

// ReadFile reads the Sitemap from file the named file.
func ReadXMLFile(name string) (*Sitemap, error) {

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	return ParseXML(data)
}

// Fetch fetches the Sitemap from url.
//
// If fetch fails, returns the status code as an error (eg.: "404").
func Fetch(url string) (*Sitemap, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}

	mediatype, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse Content-Type: %w", err)
	}

	switch mediatype {
	case "text/xml":
		return ReadXML(resp.Body)
	default:
		return nil, fmt.Errorf("unknown Content-Type: %s", mediatype)
	}

}

func CrawlHostname(hostname string, current *Sitemap) (*Sitemap, error) {

	var err error

	result := current
	if result == nil {
		result = New()
	}

	c := colly.NewCollector(
		colly.AllowedDomains(hostname),
		colly.ParseHTTPErrorResponse(),
	)

	c.OnError(func(r *colly.Response, reqErr error) {
		if r.StatusCode != 404 {
			err = fmt.Errorf("\"%s\": %w", r.Request.URL, reqErr)
		}
		result.RemoveURL(r.Request.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		result.SetURL(NewURL(r.Request.URL.String()))
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit("http://" + hostname + "/")

	if current != nil {
		for i := range current.URL {
			c.Visit(current.URL[i].Location.String())
		}
	}

	return result, err
}

func (s *Sitemap) ToXML() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	var data []byte
	var err error

	if s.isIndex {
		index := NewIndex(s.URL...)
		data, err = xml.Marshal(index)
	} else {
		urlset := NewURLSet(s.URL...)
		data, err = xml.Marshal(urlset)
	}

	if err != nil {
		return nil, err
	}

	buf := bytes.Clone([]byte(xml.Header))
	buf = append(buf, data...)

	return buf, nil
}

func (s *Sitemap) ToXMLIndent() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	var data []byte
	var err error

	if s.isIndex {
		index := NewIndex(s.URL...)
		data, err = xml.MarshalIndent(index, "", "\t")
	} else {
		urlset := NewURLSet(s.URL...)
		data, err = xml.MarshalIndent(urlset, "", "\t")
	}

	if err != nil {
		return nil, err
	}

	buf := bytes.Clone([]byte(xml.Header))
	buf = append(buf, data...)

	return buf, nil
}

func (s *Sitemap) ToTXT() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	buf := new(bytes.Buffer)

	for i := range s.URL {

		// Write Location + "\n"
		_, err := buf.WriteString(s.URL[i].Location.String() + "\n")
		if err != nil {
			return nil, fmt.Errorf("failed to write %s: %w", s.URL[i].Location, err)
		}
	}

	return buf.Bytes(), nil
}

func (s *Sitemap) IsIndex() bool {

	s.m.RLock()
	defer s.m.RUnlock()

	return s.isIndex
}

func (s *Sitemap) String() string {

	s.m.RLock()
	defer s.m.RUnlock()

	buf := new(strings.Builder)

	for i := range s.URL {
		buf.WriteString(s.URL[i].String())
		buf.WriteByte('\n')
	}

	return buf.String()
}

func (s *Sitemap) GetURL(loc string) *URL {

	if s == nil || len(loc) == 0 {
		return nil
	}

	s.m.RLock()
	defer s.m.RUnlock()

	for i := range s.URL {
		if s.URL[i].Location.Equal(Location(loc)) {
			return &s.URL[i]
		}
	}

	return nil
}

func (s *Sitemap) SetURL(u *URL) {

	if s == nil || u == nil {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	for i := range s.URL {
		if s.URL[i].Location.Equal(u.Location) {

			// Set LastMod
			if u.LastMod != nil {
				s.URL[i].LastMod = u.LastMod
			}

			// Set ChangeFreq
			if u.ChangeFreq != nil {
				s.URL[i].ChangeFreq = u.ChangeFreq
			}

			// Set Priority
			if u.Priority != nil {
				s.URL[i].Priority = u.Priority
			}

			return
		}
	}

	s.URL = append(s.URL, *u)
}

func (s *Sitemap) RemoveURL(loc string) {

	if s == nil || len(loc) == 0 {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	locIndex := -1

	for i := range s.URL {
		if s.URL[i].Location.String() == loc {
			locIndex = i
			break
		}
	}

	if locIndex == -1 {
		return
	}

	// Remove the first elem
	if locIndex == 0 {
		s.URL = s.URL[locIndex+1:]
		return
	}

	// Remove the last elem
	if locIndex == len(s.URL)-1 {
		s.URL = s.URL[:locIndex]
		return
	}

	v := s.URL[:locIndex]
	s.URL = append(v, s.URL[locIndex+1:]...)

}

func (s *Sitemap) Size() int {

	if s == nil {
		return 0
	}

	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.URL)
}

// SortByLocation sorts the URLs by Location in ascending order.
func (s *Sitemap) SortByLocation() {

	if s == nil {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	slices.SortStableFunc(s.URL, func(a, b URL) int {
		return strings.Compare(a.Location.String(), b.Location.String())
	})
}
