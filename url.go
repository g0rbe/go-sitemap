package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
)

var (
	ErrEmptyLoc = errors.New("empty loc")
)

type URL struct {
	Loc        Location         `xml:"loc"`
	LastMod    LastModification `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFrequency  `xml:"changefreq,omitempty"`
	Priority   Priority         `xml:"priority,omitempty"`
}

func (u URL) String() string {
	return fmt.Sprintf("%s\t%s\t%s\t%s", u.Loc.String(), u.LastMod.String(), u.ChangeFreq.String(), u.Priority.String())
}

func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		Loc        Location         `xml:"loc"`
		LastMod    LastModification `xml:"lastmod,omitempty"`
		ChangeFreq ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority   Priority         `xml:"priority,omitempty"`
	}{}

	err := d.DecodeElement(&v, &start)

	if err != nil {
		return err
	}

	if v.Loc.IsEmpty() {
		return ErrEmptyLoc
	}

	*u = *(*URL)(&v)

	return nil
}

func (u URL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "url"

	v := struct {
		Loc        Location         `xml:"loc"`
		LastMod    LastModification `xml:"lastmod,omitempty"`
		ChangeFreq ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority   Priority         `xml:"priority,omitempty"`
	}{
		Loc:        u.Loc,
		LastMod:    u.LastMod,
		ChangeFreq: u.ChangeFreq,
		Priority:   u.Priority,
	}

	return e.EncodeElement(v, start)

}
