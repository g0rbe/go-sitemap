package sitemap

import "encoding/xml"

// Index encapsulates information about all of the Sitemaps in the file.
//
//	<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//	  <sitemap>
//	    <loc>http://www.example.com/sitemap1.xml.gz</loc>
//	    <lastmod>2004-10-01T18:23:17+00:00</lastmod>
//	  </sitemap>
//	  <sitemap>
//	    <loc>http://www.example.com/sitemap2.xml.gz</loc>
//	    <lastmod>2005-01-01</lastmod>
//	  </sitemap>
//	</sitemapindex>
type Index []URL

func (i Index) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Add xml.Header before encoding
	err := e.EncodeToken(xml.ProcInst{Target: "xml", Inst: []byte("version=\"1.0\" encoding=\"UTF-8\"")})
	if err != nil {
		return err
	}

	// Set the first token name to "sitemapindex"
	if start.Name.Local != "sitemapindex" {
		start.Name.Local = "sitemapindex"
	}

	// Set "xmlns" attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: XMLNameSpace})

	v := struct {
		URL []URL `xml:"sitemap"`
	}{
		URL: i,
	}

	return e.EncodeElement(v, start)
}

func (i *Index) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		URL []URL `xml:"sitemap"`
	}{}

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	*i = v.URL

	return nil
}
