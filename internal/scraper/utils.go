package scraper

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// LoadUserAgents loads UAs from txt file
func LoadUserAgents(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open user_agents file: %v", err)
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
		log.Fatalf("error reading user_agents file: %v", err)
	}
	if len(agents) == 0 {
		log.Fatalf("user_agents.txt is empty")
	}
	return agents
}

// PickRandomUA selects one UA randomly
func PickRandomUA(agents []string) string {
	rand.Seed(time.Now().UnixNano())
	return agents[rand.Intn(len(agents))]
}
