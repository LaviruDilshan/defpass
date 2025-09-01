package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/LaviruDilshan/defpass/internal/scraper"
	"github.com/LaviruDilshan/defpass/internal/ui"
)

func main() {
	ui.PrintBanner()

	agents := scraper.LoadUserAgents()

	vendors, err := scraper.LoadVendors(agents)
	if err != nil {
		log.Fatalf("failed to load vendors: %v", err)
	}
	if len(vendors) == 0 {
		fmt.Println("No vendors found.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nAvailable Vendors:")
		for i, vendor := range vendors {
			fmt.Printf("%d. %s\n", i+1, vendor)
		}

		fmt.Print("\nEnter vendor number (or 'q' to quit): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) == "q" {
			break
		}

		var idx int
		_, err := fmt.Sscanf(choice, "%d", &idx)
		if err != nil || idx < 1 || idx > len(vendors) {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		vendor := vendors[idx-1]
		fmt.Printf("\nFetching default passwords for vendor: %s...\n\n", vendor)

		results, err := scraper.LoadVendorPasswords(vendor, agents)
		if err != nil {
			log.Println("Error fetching passwords:", err)
			continue
		}

		if len(results) == 0 {
			fmt.Println("No passwords found for this vendor.")
		} else {
			for i, entry := range results {
				fmt.Printf("---- Entry %d ----\n", i+1)
				for k, v := range entry {
					fmt.Printf("%s: %s\n", k, v)
				}
				fmt.Println("------------------")
			}
		}

		fmt.Print("\nPress Enter to return to vendor list, or 'q' to quit: ")
		back, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(back)) == "q" {
			break
		}
	}
}
