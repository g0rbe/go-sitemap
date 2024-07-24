package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
)

var (
	ErrNotURLSet = errors.New("not urlset")
)

type Sitemap struct {
	urls []URL
}

func NewSitemap(urls []URL) *Sitemap {
	v := new(Sitemap)

	if len(urls) > 0 {
		v.urls = append(v.urls, urls...)
	}

	return v
}

func FetchSitemap(url string) (*Sitemap, error) {

	data, err := download(url)
	if err != nil {
		return nil, err
	}

	v := new(Sitemap)

	err = xml.Unmarshal(data, v)

	return v, err
}

func FetchSitemaps(urls []string) (*Sitemap, error) {

	s := new(Sitemap)

	for i := range urls {

		v, err := FetchSitemap(urls[i])
		if err != nil {
			return s, fmt.Errorf("failed to fetch \"%s\": %w", urls[i], err)
		}

		s.AddPages(v.urls)
	}

	return s, nil
}

func Fetch(url string) (*Sitemap, error) {

	data, err := download(url)
	if err != nil {
		return nil, err
	}

	format, err := GetFormat(data)
	if err != nil {
		return nil, fmt.Errorf("failed to get format: %w", err)
	}

	switch format {
	case FormatURLSet:

		v := new(Sitemap)
		err = xml.Unmarshal(data, v)
		return v, err

	case FormatIndex:

		v := new(Index)
		err = xml.Unmarshal(data, v)
		if err != nil {
			return nil, err
		}
		return FetchSitemaps(v.Locations())

	default:

		return nil, fmt.Errorf("invalid format: %s", format)
	}
}

func (s *Sitemap) NumPages() int {
	return len(s.urls)
}

func (s *Sitemap) Pages() []URL {
	return s.urls
}

func (s *Sitemap) AddPages(u []URL) {
	s.urls = append(s.urls, u...)
}

func (s *Sitemap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		URLs []URL `xml:"url"`
	}{}

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	s.urls = v.URLs

	return nil
}

func (s *Sitemap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "urlset"

	v := struct {
		URLs []URL `xml:"url"`
	}{
		URLs: s.urls,
	}

	return e.EncodeElement(v, start)
}
