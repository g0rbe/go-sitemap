package sitemap

const XMLNameSpace = "http://www.sitemaps.org/schemas/sitemap/0.9"

type Sitemap interface {
	ToXML() ([]byte, error)
}
