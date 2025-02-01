package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"slices"
	"strings"
	"sync"
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

func New() Sitemap {
	return Sitemap{m: new(sync.RWMutex)}
}

func (s Sitemap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	var v any

	if s.isIndex {
		v = Index(s.URL)
	} else {
		v = URLSet(s.URL)
	}

	return e.EncodeElement(v, start)
}

func (s *Sitemap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	switch start.Name.Local {
	case "sitemapindex":

		v := Index{}

		err := v.UnmarshalXML(d, start)
		if err != nil {
			return err
		}

		s.URL = []URL(v)

		return nil

	case "urlset":

		v := URLSet{}
		err := v.UnmarshalXML(d, start)
		if err != nil {
			return err
		}

		s.URL = []URL(v)

		return nil

	default:
		return fmt.Errorf("unknown start element: %s", start.Name.Local)
	}
}

func (s *Sitemap) SetIndex(v bool) {

	s.m.Lock()
	defer s.m.Unlock()

	s.isIndex = v
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
		if s.URL[i].Location == loc {
			return &s.URL[i]
		}
	}

	return nil
}

// AppendURL append URL u if unique.
//
// Returns true if URL u is unique and appended to the Sitemap.
func (s *Sitemap) AppendURL(u URL) bool {

	if s.GetURL(u.Location) != nil {
		return false
	}

	s.m.Lock()
	defer s.m.Unlock()

	s.URL = append(s.URL, u)

	return true
}

// RemoveURL removes the URL with the given Location loc from Sitemap s.
//
// Returns true if the URL with Location loc is removed from Sitemap s.
func (s *Sitemap) RemoveURL(loc string) bool {

	if s == nil || len(loc) == 0 {
		return false
	}

	s.m.Lock()
	defer s.m.Unlock()

	locIndex := -1

	for i := range s.URL {
		if s.URL[i].Location == loc {
			locIndex = i
			break
		}
	}

	if locIndex == -1 {
		return false
	}

	// Remove the first elem
	if locIndex == 0 {
		s.URL = s.URL[locIndex+1:]
		return true
	}

	// Remove the last elem
	if locIndex == len(s.URL)-1 {
		s.URL = s.URL[:locIndex]
		return true
	}

	v := s.URL[:locIndex]
	s.URL = append(v, s.URL[locIndex+1:]...)

	return true
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
		return strings.Compare(a.Location, b.Location)
	})
}
