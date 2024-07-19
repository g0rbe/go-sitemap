package sitemap

import "fmt"

type URL struct {
	Loc        Location  `xml:"loc"`
	LastMod    Time      `xml:"lastmod"`
	ChangeFreq Frequency `xml:"changefreq"`
	Priority   Priority  `xml:"priority"`
}

func (u *URL) String() string {
	return fmt.Sprintf("%s %s %s %s", u.Loc.String(), u.LastMod.String(), u.ChangeFreq.String(), u.Priority.String())
}
