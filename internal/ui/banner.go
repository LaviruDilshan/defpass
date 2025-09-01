package ui

import "fmt"

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

func PrintBanner() {
	banner := `
     ____        __    ____                  
    |  _ \  ___ / _|  |  _ \ __ _ ___ ___ ___ 
    | | | |/ _ \ |_   | |_) / _` + "`" + ` / __/ __/ _ \
    | |_| |  __/  _|  |  __/ (_| \__ \__ \  __/
    |____/ \___|_|    |_|   \__,_|___/___/\___|
    `
	// Print banner in cyan
	fmt.Println(Cyan + banner + Reset)

	// Print tool name in cyan
	fmt.Println(Cyan + "           Default Password Finder" + Reset)

	// Author and socials with colors
	fmt.Println(Green + " Author   : Laviru Dilshan" + Reset)
	fmt.Println(Yellow + " LinkedIn : https://linkedin.com/in/lavirudilshan" + Reset)
	fmt.Println(Yellow + " X        : https://x.com/LaviruDilshan" + Reset)
	fmt.Println(Yellow + " Instagram: https://instagram.com/lavirudilshn" + Reset)
	fmt.Println(Cyan + "--------------------------------------------------" + Reset)
}
