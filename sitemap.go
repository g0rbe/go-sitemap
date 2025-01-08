package sitemap

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	URLs    []URL
	isIndex bool
	m       *sync.RWMutex
}

func New() *Sitemap {
	s := new(Sitemap)
	s.m = new(sync.RWMutex)
	return s
}

// Parse reads the Sitemap from data.
func Parse(data []byte) (*Sitemap, error) {

	s := New()

	if IsIndex(data) {
		i, err := ParseIndex(data)
		if err != nil {
			return nil, fmt.Errorf("filed to parse Index: %w", err)
		}
		s.isIndex = true
		s.URLs = i.Sitemaps
	} else {
		u, err := ParseURLSet(data)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URLSet: %w", err)
		}

		s.URLs = u.URLs
	}

	return s, nil
}

// Read reads the Sitemap from r.
func Read(r io.Reader) (*Sitemap, error) {

	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	return Parse(data)
}

// ReadFile reads the Sitemap from file the named file.
func ReadFile(name string) (*Sitemap, error) {

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	return Parse(data)
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

	return Read(resp.Body)
}

func Crawler(loc string, current *Sitemap) (*Sitemap, error) {

	target, err := url.Parse(loc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", loc, err)
	}

	result := current
	if result == nil {
		result = New()
	}

	c := colly.NewCollector(
		colly.AllowedDomains(target.Hostname()),
		colly.ParseHTTPErrorResponse(),
	)

	c.OnError(func(r *colly.Response, reqErr error) {
		if r.StatusCode != 404 && err != nil {
			err = fmt.Errorf("\"%s\": %w", r.Request.URL, reqErr)
		}
		result.RemoveURL(r.Request.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		result.SetURL(NewURL(NewLocation(r.Request.URL.String())))
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit(target.String())

	if current != nil {
		for i := range current.URLs {
			c.Visit(current.URLs[i].Location.String())
		}
	}

	return result, err
}

func (s *Sitemap) ToXML() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	if s.isIndex {
		i := new(Index)
		i.Sitemaps = s.URLs
		return i.ToXML()
	}

	u := new(URLSet)
	u.URLs = s.URLs
	return u.ToXML()
}

func (s *Sitemap) ToXMLIndent() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	if s.isIndex {
		i := new(Index)
		i.Sitemaps = s.URLs
		return i.ToXMLIndent()
	}

	u := new(URLSet)
	u.URLs = s.URLs
	return u.ToXMLIndent()
}

func (s *Sitemap) ToTXT() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	buf := new(bytes.Buffer)

	for i := range s.URLs {

		// Write Location + "\n"
		_, err := buf.WriteString(s.URLs[i].Location.String() + "\n")
		if err != nil {
			return nil, fmt.Errorf("failed to write %s: %w", s.URLs[i].Location, err)
		}
	}

	return buf.Bytes(), nil
}

func (s *Sitemap) ToJSON() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	if s.isIndex {
		i := new(Index)
		i.Sitemaps = s.URLs
		return i.ToJSON()
	}

	u := new(URLSet)
	u.URLs = s.URLs
	return u.ToJSON()
}

func (s *Sitemap) ToJSONIndent() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	if s.isIndex {
		i := new(Index)
		i.Sitemaps = s.URLs
		return i.ToJSONIndent()
	}

	u := new(URLSet)
	u.URLs = s.URLs
	return u.ToJSONIndent()
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

	for i := range s.URLs {
		buf.WriteString(s.URLs[i].String())
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

	for i := range s.URLs {
		if s.URLs[i].Location.Equal(NewLocation(loc)) {
			return &s.URLs[i]
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

	for i := range s.URLs {
		if s.URLs[i].Location.Equal(u.Location) {

			// Set LastMod
			if u.LastMod != nil {
				s.URLs[i].LastMod = u.LastMod
			}

			// Set ChangeFreq
			if u.ChangeFreq != nil {
				s.URLs[i].ChangeFreq = u.ChangeFreq
			}

			// Set Priority
			if u.Priority != nil {
				s.URLs[i].Priority = u.Priority
			}

			// Set Comment
			if u.Comment != nil {
				s.URLs[i].Comment = u.Comment
			}

			return
		}
	}

	s.URLs = append(s.URLs, *u)
}

func (s *Sitemap) RemoveURL(loc string) {

	if s == nil || len(loc) == 0 {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	locIndex := -1

	for i := range s.URLs {
		if s.URLs[i].Location.String() == loc {
			locIndex = i
			break
		}
	}

	if locIndex == -1 {
		return
	}

	// Remove the first elem
	if locIndex == 0 {
		s.URLs = s.URLs[locIndex+1:]
		return
	}

	// Remove the last elem
	if locIndex == len(s.URLs)-1 {
		s.URLs = s.URLs[:locIndex]
		return
	}

	v := s.URLs[:locIndex]
	s.URLs = append(v, s.URLs[locIndex+1:]...)

}

func (s *Sitemap) Size() int {

	if s == nil {
		return 0
	}

	s.m.RLock()
	defer s.m.RUnlock()

	return len(s.URLs)
}

// SortByLocation sorts the URLs by Location in ascending order.
func (s *Sitemap) SortByLocation() {

	if s == nil {
		return
	}

	s.m.Lock()
	defer s.m.Unlock()

	slices.SortStableFunc(s.URLs, func(a, b URL) int {
		return strings.Compare(a.Location.String(), b.Location.String())
	})
}
