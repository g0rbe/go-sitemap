package sitemap

import (
	"encoding/xml"
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
type URLSet []URL

func (s URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Add xml.Header before encoding
	err := e.EncodeToken(xml.ProcInst{Target: "xml", Inst: []byte("version=\"1.0\" encoding=\"UTF-8\"")})
	if err != nil {
		return err
	}

	// Set the first token name to "sitemapindex"
	if start.Name.Local != "urlset" {
		start.Name.Local = "urlset"
	}

	// Set "xmlns" attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: XMLNameSpace})

	v := struct {
		URL []URL `xml:"url"`
	}{
		URL: s,
	}

	return e.EncodeElement(v, start)
}

func (s *URLSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		URL []URL `xml:"url"`
	}{}

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	*s = v.URL

	return nil
}
