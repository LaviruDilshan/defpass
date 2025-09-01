package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Entry map[string]string

// LoadVendorPasswords scrapes vendor default passwords
func LoadVendorPasswords(vendor string, agents []string) ([]Entry, error) {
	var results []Entry

	c := colly.NewCollector()

	// Randomize UA for each request
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", PickRandomUA(agents))
	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
		entry := Entry{}
		e.ForEach("tr", func(_ int, tr *colly.HTMLElement) {
			tds := []string{}
			tr.ForEach("td", func(_ int, td *colly.HTMLElement) {
				tds = append(tds, strings.TrimSpace(td.Text))
			})
			if len(tds) == 2 {
				entry[tds[0]] = tds[1]
			}
		})
		if len(entry) > 0 {
			results = append(results, entry)
		}
	})

	err := c.Visit(baseURL + "?vendor=" + vendor)
	if err != nil {
		log.Println("Error fetching vendor passwords:", err)
	}

	return results, nil
}
