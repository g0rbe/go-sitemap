package sitemap

import (
	"encoding/xml"
)

type Index struct {
	sitemaps []URL
}

func NewIndex(urls []URL) *Index {

	v := new(Index)

	if len(urls) > 0 {
		v.sitemaps = append(v.sitemaps, urls...)
	}

	return v
}

func FetchIndex(url string) (*Index, error) {

	data, err := download(url)
	if err != nil {
		return nil, err
	}

	v := new(Index)

	err = xml.Unmarshal(data, v)

	return v, err
}

func (i *Index) URLs() []URL {
	return i.sitemaps
}

func (i *Index) Sitemaps() []string {

	v := make([]string, 0, len(i.sitemaps))

	for j := range i.sitemaps {
		v = append(v, i.sitemaps[j].Loc.String())
	}

	return v
}

func (i *Index) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		URLs []URL `xml:"sitemap"`
	}{}

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	i.sitemaps = v.URLs

	return nil
}

func (i *Index) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "sitemapindex"

	v := struct {
		Sitemap []URL `xml:"sitemap"`
	}{
		Sitemap: i.sitemaps,
	}

	return e.EncodeElement(v, start)
}
