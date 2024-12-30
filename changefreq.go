package sitemap

import (
	"encoding/xml"
	"strings"
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
type ChangeFrequency string

// Valid values for ChangeFrequency
var (
	ChangeFrequencyAlways  ChangeFrequency = "always"
	ChangeFrequencyHourly  ChangeFrequency = "hourly"
	ChangeFrequencyDaily   ChangeFrequency = "daily"
	ChangeFrequencyWeekly  ChangeFrequency = "weekly"
	ChangeFrequencyMonthly ChangeFrequency = "monthly"
	ChangeFrequencyYearly  ChangeFrequency = "yearly"
	ChangeFrequencyNever   ChangeFrequency = "never"
)

// ParseChangeFrequency parses ChangeFrequency from v.
//
// If v is not a valid value for ChangeFrequency, returns nil.
func ParseChangeFrequency(v string) *ChangeFrequency {

	switch strings.ToLower(v) {
	case "always":
		return &ChangeFrequencyAlways
	case "hourly":
		return &ChangeFrequencyHourly
	case "daily":
		return &ChangeFrequencyDaily
	case "weekly":
		return &ChangeFrequencyWeekly
	case "monthly":
		return &ChangeFrequencyMonthly
	case "yearly":
		return &ChangeFrequencyYearly
	case "never":
		return &ChangeFrequencyNever
	default:
		return nil
	}
}

func (c *ChangeFrequency) String() string {
	return string(*c)
}

func (c *ChangeFrequency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Change the start and end tag to "changefreq"
	if start.Name.Local != "changefreq" {
		start.Name.Local = "changefreq"
	}

	return e.EncodeElement(c.String(), start)
}
