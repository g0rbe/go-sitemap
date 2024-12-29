package sitemap

import "encoding/xml"

// The priority of this URL relative to other URLs on your site. Valid values range from 0.0 to 1.0.
// This value does not affect how your pages are compared to pages on other sites—it only lets the search engines know which pages you deem most important for the crawlers.
//
// The default priority of a page is 0.5.
//
// Please note that the priority you assign to a page is not likely to influence the position of your URLs in a search engine's result pages.
// Search engines may use this information when selecting between URLs on the same site, so you can use this tag to increase the likelihood that your most important pages are present in a search index.
//
// Also, please note that assigning a high priority to all of the URLs on your site is not likely to help you.
// Since the priority is relative, it is only used to select between URLs on your site.
//
// Example:
//
//	<priority>0.8</priority>
type Priority string

// NewPriority returns a new Priority with the given value v.
// If v is an emty string(""), returns nil.
//
// This function sets the XMLName to "priority".
func NewPriority(v string) *Priority {
	if v == "" {
		return nil
	}
	return (*Priority)(&v)
}

func (p *Priority) String() string {
	return string(*p)
}

func (p *Priority) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := new(string)

	err := d.DecodeElement(v, &start)
	if err != nil {
		return err
	}

	*p = Priority(*v)

	return nil
}

func (p *Priority) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "priority"
	if start.Name.Local != "priority" {
		start.Name.Local = "priority"
	}

	return e.EncodeElement(p.String(), start)
}
