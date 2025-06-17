package webserver

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/csmith/splendid/backend/model"
)

func Redact(value any, playerID model.PlayerID) ([]byte, error) {
	// Marshal the value to JSON
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	result := string(jsonBytes)

	// Find all redacted segments: ["REDACT-...", "player1,player2", ...
	startRegex := regexp.MustCompile(`\["(REDACT-[^"]+)","([^"]*)"`)
	matches := startRegex.FindAllStringSubmatch(result, -1)

	for _, match := range matches {
		if len(match) != 3 {
			continue
		}

		token := match[1]
		playerIDsStr := match[2]

		// Build a regex to match the entire redacted segment with this specific token
		fullRegex := regexp.MustCompile(`\["` + regexp.QuoteMeta(token) + `","` + regexp.QuoteMeta(playerIDsStr) + `",(.*?),"` + regexp.QuoteMeta(token) + `"\]`)

		// Replace this specific segment
		result = fullRegex.ReplaceAllStringFunc(result, func(fullMatch string) string {
			fullMatches := fullRegex.FindStringSubmatch(fullMatch)
			if len(fullMatches) != 2 {
				return `"REDACTED-ERROR"` // Return error on safe side
			}

			// Split player IDs to avoid substring matches
			playerIDs := strings.Split(playerIDsStr, ",")

			// Check if the player ID is in the visibility list
			for _, id := range playerIDs {
				if strings.TrimSpace(id) == string(playerID) {
					// Player can see the value, return the actual value
					return fullMatches[1]
				}
			}

			// Player cannot see the value, return "REDACTED"
			return `"REDACTED"`
		})
	}

	return []byte(result), nil
}
