package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"net/url"
	"os"

	"git.gorbe.io/go/sitemap"
	"github.com/gocolly/colly"
)

var (
	urlFlag     string
	urlParsed   *url.URL
	currentFlag string
	format      string
	outName     string
	outFile     *os.File
)

func init() {

	var err error

	flag.StringVar(&urlFlag, "url", "", "URL to crawl")
	flag.StringVar(&currentFlag, "current", "", "URL of the current sitemap to update")
	flag.StringVar(&format, "format", "xml", "Format out the output sitemap (\"xml\" / \"xmli\" / \"txt\")")
	flag.StringVar(&outName, "out", "/dev/stdout", "Name of output file")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: sitemap [ -url <url> ] [ -current <url> ] [ -format <format> ] [ -out <name> ]\n\n")
		fmt.Printf("Flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Check URL
	if len(urlFlag) == 0 {
		fmt.Fprintf(os.Stderr, "Fail: \"-url\" falg is missing!\n")
		os.Exit(1)
	}

	urlParsed, err = url.Parse(urlFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail: failed to parse URL \"%s\": %s\n", urlFlag, err)
		os.Exit(1)
	}

	// Check format
	if format != "xml" && format != "xmli" && format != "txt" {
		fmt.Fprintf(os.Stderr, "Fail: invalid format: \"%s\"\n", format)
		os.Exit(1)
	}

	// Check output name
	if len(outName) == 0 {
		fmt.Fprintf(os.Stderr, "Fail: the \"-out\" flag is empty!\n")
		os.Exit(1)
	}

}

func CrawlHostname(u *url.URL, sm *sitemap.Sitemap) error {

	var err error

	c := colly.NewCollector(
		colly.AllowedDomains(u.Hostname()),
		colly.ParseHTTPErrorResponse(),
	)

	c.OnError(func(r *colly.Response, reqErr error) {
		err = fmt.Errorf("\"%s\": %w", r.Request.URL, reqErr)

		fmt.Fprintf(os.Stderr, "[!] Crawler error: %s\n", err)

		sm.RemoveURL(r.Request.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		if sm.AppendURL(sitemap.URL{Location: r.Request.URL.String()}) {
			fmt.Fprintf(os.Stderr, "[+] New location found: %s\n", r.Request.URL.String())
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.Visit(u.String())

	for i := range sm.URL {
		c.Visit(sm.URL[i].Location)
	}

	return err
}

func main() {

	var sm = sitemap.NewSitemap()

	if len(currentFlag) > 0 {
		currentSm, err := sitemap.Fetch(currentFlag)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail: failed to fecth current sitemap \"%s\": %s\n", currentFlag, err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "[i] Fetched %d location from \"%s\"\n", currentSm.Size(), currentFlag)

		sm = currentSm
	}

	err := CrawlHostname(urlParsed, &sm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "[i] Found %d location on \"%s\"\n", sm.Size(), urlFlag)

	var buf []byte

	switch format {
	case "xml":
		buf, err = xml.Marshal(sm)
	case "xmli":
		buf, err = xml.MarshalIndent(sm, "", "\t")
	case "txt":
		buf = []byte(sm.String())
	default:
		fmt.Fprintf(os.Stderr, "Invalid format: %s\n", format)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail: failed to format sitemap: %s\n", err)
		os.Exit(1)
	}

	outFile, err = os.OpenFile(outName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail: failed to open \"%s\": %s\n", outName, err)
		os.Exit(1)
	}

	fmt.Fprintf(outFile, "%s", buf)

}
