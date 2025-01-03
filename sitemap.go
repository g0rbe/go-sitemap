package sitemap

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const XMLNameSpace = "http://www.sitemaps.org/schemas/sitemap/0.9"

func IsIndex(data []byte) bool {
	return bytes.Contains(data, []byte("sitemapindex"))
}

type Sitemap struct {
	URLs  []URL
	index bool
	m     *sync.RWMutex
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
		s.index = true
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

// Fetch fetches the Sitemap from url.
func Fetch(url string) (*Sitemap, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return Read(resp.Body)
}

func (s *Sitemap) ToXML() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	if s.index {
		i := new(Index)
		i.Sitemaps = s.URLs
		return i.ToXML()
	}

	u := new(URLSet)
	u.URLs = s.URLs
	return u.ToXML()
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
