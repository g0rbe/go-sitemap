package sitemap

import (
	"bytes"
)

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
// Example:
//
//	<changefreq>monthly</changefreq>
type ChangeFrequency []byte

// Valid values for ChangeFrequency
var (
	ChangeFrequencyAlways  ChangeFrequency = []byte("always")
	ChangeFrequencyHourly  ChangeFrequency = []byte("hourly")
	ChangeFrequencyDaily   ChangeFrequency = []byte("daily")
	ChangeFrequencyWeekly  ChangeFrequency = []byte("weekly")
	ChangeFrequencyMonthly ChangeFrequency = []byte("monthly")
	ChangeFrequencyYearly  ChangeFrequency = []byte("yearly")
	ChangeFrequencyNever   ChangeFrequency = []byte("never")
)

// ParseChangeFrequency parses ChangeFrequency from v.
//
// If v is not a valid value for ChangeFrequency, returns nil.
func (c ChangeFrequency) IsValid() bool {

	switch v := bytes.ToLower(c); true {
	case bytes.Equal(v, ChangeFrequencyAlways):
		return true
	case bytes.Equal(v, ChangeFrequencyHourly):
		return true
	case bytes.Equal(v, ChangeFrequencyDaily):
		return true
	case bytes.Equal(v, ChangeFrequencyWeekly):
		return true
	case bytes.Equal(v, ChangeFrequencyMonthly):
		return true
	case bytes.Equal(v, ChangeFrequencyYearly):
		return true
	case bytes.Equal(v, ChangeFrequencyNever):
		return true
	default:
		return false
	}
}

func (c ChangeFrequency) String() string {
	return string(c)
}

// func (c *ChangeFrequency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

// 	// Change the start and end tag to "changefreq"
// 	if start.Name.Local != "changefreq" {
// 		start.Name.Local = "changefreq"
// 	}

// 	return e.EncodeElement(c.String(), start)
// }
