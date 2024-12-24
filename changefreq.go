package sitemap

import "encoding/xml"

// How frequently the page is likely to change. This value provides general information to search engines and may not correlate exactly to how often they crawl the page.
//
// Valid values are:
//   - always
//   - hourly
//   - daily
//   - weekly
//   - monthly
//   - yearly
//   - never
//
// The value "always" should be used to describe documents that change each time they are accessed. The value "never" should be used to describe archived URLs.
//
// Please note that the value of this tag is considered a hint and not a command.
// Even though search engine crawlers may consider this information when making decisions, they may crawl pages marked "hourly" less frequently than that, and they may crawl pages marked "yearly" more frequently than that.
// Crawlers may periodically crawl pages marked "never" so that they can handle unexpected changes to those pages.
//
//	<changefreq>monthly</changefreq>
type ChangeFreq struct {
	XMLName xml.Name `xml:"changefreq"`
	Value   []byte   `xml:",chardata"`
}

// NewChangeFreq returns a new ChangeFreq with the given value v.
// If v is an emty string(""), returns nil.
//
// This function sets the XMLName to "changefreq".
func NewChangeFreq(v string) *ChangeFreq {
	if v == "" {
		return nil
	}
	return &ChangeFreq{XMLName: xml.Name{Local: "changefreq"}, Value: []byte(v)}
}

func (c *ChangeFreq) String() string {
	return string(c.Value)
}
