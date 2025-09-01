package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

const baseURL = "https://cirt.net/passwords"

// LoadVendors scrapes vendors, skips google_ad_client entries
func LoadVendors(agents []string) ([]string, error) {
	var vendors []string
	c := colly.NewCollector()

	// Randomize UA for each request
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", PickRandomUA(agents))
	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			row.ForEach("td, th", func(_ int, col *colly.HTMLElement) {
				text := strings.TrimSpace(col.Text)
				if text != "" && !strings.Contains(strings.ToLower(text), "google_ad_client") {
					if !contains(vendors, text) {
						vendors = append(vendors, text)
					}
				}
			})
		})
	})

	err := c.Visit(baseURL)
	if err != nil {
		return nil, err
	}

	return vendors, nil
}

// contains checks if slice has a value
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
