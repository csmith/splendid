package model

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"
)

type Redactable[T any] struct {
	value     T
	visibleTo map[PlayerID]bool
}

func NewRedactable[T any](value T, playerIDs ...PlayerID) Redactable[T] {
	visibleTo := make(map[PlayerID]bool)
	for _, playerID := range playerIDs {
		visibleTo[playerID] = true
	}
	return Redactable[T]{
		value:     value,
		visibleTo: visibleTo,
	}
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
	for playerID, visible := range r.visibleTo {
		if visible {
			if visibility.Len() > 0 {
				visibility.WriteString(",")
			}
			visibility.WriteString(string(playerID))
		}
	}
	result = append(result, visibility.String())

	// Add the actual value as raw JSON
	valueBytes, err := json.Marshal(&r.value)
	if err != nil {
		return nil, err
	}
	var rawValue json.RawMessage = valueBytes
	result = append(result, rawValue)

	// Add same redact token at the end
	result = append(result, token)

	return json.Marshal(result)
}

func (r Redactable[T]) Value() T {
	return r.value
}

func (r Redactable[T]) WithVisibility(playerIDs ...PlayerID) Redactable[T] {
	return NewRedactable(r.value, playerIDs...)
}

func generateRedactToken() (string, error) {
	bytes := make([]byte, 8) // 16 hex characters
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return "REDACT-" + hex.EncodeToString(bytes), nil
}
