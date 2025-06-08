package model

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type Redactable[T any] struct {
	Value     T
	VisibleTo map[PlayerID]bool
}

func (r Redactable[T]) MarshalJSON() ([]byte, error) {
	result := make([]interface{}, 0)

	// Generate redact token (same for start and end)
	token, err := generateRedactToken()
	if err != nil {
		return nil, err
	}
	result = append(result, token)

	// Add player IDs who can see the value
	visibility := strings.Builder{}
	for playerID, visible := range r.VisibleTo {
		if visible {
			if visibility.Len() > 0 {
				visibility.WriteString(",")
			}
			visibility.WriteString(string(playerID))
		}
	}
	result = append(result, visibility.String())

	// Add the actual value as raw JSON
	valueBytes, err := json.Marshal(r.Value)
	if err != nil {
		return nil, err
	}
	var rawValue json.RawMessage = valueBytes
	result = append(result, rawValue)

	// Add same redact token at the end
	result = append(result, token)

	return json.Marshal(result)
}

func generateRedactToken() (string, error) {
	bytes := make([]byte, 8) // 16 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return "REDACT-" + hex.EncodeToString(bytes), nil
}
