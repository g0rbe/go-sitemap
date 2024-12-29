package sitemap

import "bytes"

const XMLNameSpace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Sitemap interface {
	ToXML() ([]byte, error)
}

func SitemapIsIndex(data []byte) bool {
	return bytes.Contains(data, []byte("sitemapindex"))
}
