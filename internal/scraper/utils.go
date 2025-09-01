package scraper

import (
	"bufio"
	"embed"
	"log"
	"math/rand"
	"strings"
	"time"
)

//go:embed user_agents.txt
var uaFile embed.FS

// LoadUserAgents loads UAs from embedded txt file
func LoadUserAgents() []string {
	file, err := uaFile.Open("user_agents.txt")
	if err != nil {
		log.Fatalf("failed to open embedded user_agents file: %v", err)
	}
	defer file.Close()

	var agents []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			agents = append(agents, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading embedded user_agents file: %v", err)
	}
	if len(agents) == 0 {
		log.Fatalf("embedded user_agents.txt is empty")
	}
	return agents
}

// PickRandomUA selects one UA randomly
func PickRandomUA(agents []string) string {
	rand.Seed(time.Now().UnixNano())
	return agents[rand.Intn(len(agents))]
}
