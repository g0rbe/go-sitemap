package sitemap

import "strconv"

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
type Priority []byte

func (p Priority) Float64() (float64, error) {
	return strconv.ParseFloat(p.String(), 64)
}

func (p Priority) String() string {
	return string(p)
}
