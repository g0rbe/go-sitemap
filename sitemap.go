package sitemap

type Sitemap interface {
	ToXML() ([]byte, error)
}
