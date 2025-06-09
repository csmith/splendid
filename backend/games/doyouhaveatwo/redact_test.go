package doyouhaveatwo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/csmith/splendid/backend/games/doyouhaveatwo/model"
)

func TestRedact_BasicVisiblePlayer(t *testing.T) {
	redactable := model.Redactable[string]{
		Value: "secret-data",
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
			"player2": false,
		},
	}

	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"secret-data"`, string(result))
}

func TestRedact_CardVisiblePlayer(t *testing.T) {
	redactable := model.Redactable[model.Card]{
		Value: model.CardPrincess,
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
			"player2": false,
		},
	}

	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"Princess"`, string(result))
}

func TestRedact_BasicNonVisiblePlayer(t *testing.T) {
	redactable := model.Redactable[string]{
		Value: "secret-data",
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
			"player2": false,
		},
	}

	result, err := Redact(redactable, "player2")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}

func TestRedact_UnknownPlayer(t *testing.T) {
	redactable := model.Redactable[string]{
		Value: "secret-data",
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
			"player2": false,
		},
	}

	result, err := Redact(redactable, "player3")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}

func TestRedact_NestedRedactable(t *testing.T) {
	innerRedactable := model.Redactable[string]{
		Value: "inner-secret",
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
		},
	}

	outerRedactable := model.Redactable[model.Redactable[string]]{
		Value: innerRedactable,
		VisibleTo: map[model.PlayerID]bool{
			"player1": true,
			"player2": true,
		},
	}

	// Test with player1 - can see both outer and inner
	result1, err := Redact(outerRedactable, "player1")
	require.NoError(t, err)

	expected1 := `"inner-secret"`
	assert.Equal(t, expected1, string(result1))

	// Test with player2 - can see outer but not inner
	result2, err := Redact(outerRedactable, "player2")
	require.NoError(t, err)

	expected2 := `"REDACTED"`
	assert.Equal(t, expected2, string(result2))
}

func TestRedact_MultipleRedactables(t *testing.T) {
	type TestStruct struct {
		Public  string                   `json:"public"`
		Secret1 model.Redactable[string] `json:"secret1"`
		Secret2 model.Redactable[int]    `json:"secret2"`
	}

	data := TestStruct{
		Public: "everyone-can-see",
		Secret1: model.Redactable[string]{
			Value: "secret-string",
			VisibleTo: map[model.PlayerID]bool{
				"player1": true,
			},
		},
		Secret2: model.Redactable[int]{
			Value: 42,
			VisibleTo: map[model.PlayerID]bool{
				"player2": true,
			},
		},
	}

	// Test with player1 - can see secret1 but not secret2
	result1, err := Redact(data, "player1")
	require.NoError(t, err)

	expected1 := `{"public":"everyone-can-see","secret1":"secret-string","secret2":"REDACTED"}`
	assert.Equal(t, expected1, string(result1))

	// Test with player2 - can see secret2 but not secret1
	result2, err := Redact(data, "player2")
	require.NoError(t, err)

	expected2 := `{"public":"everyone-can-see","secret1":"REDACTED","secret2":42}`
	assert.Equal(t, expected2, string(result2))
}

func TestRedact_EmptyVisibility(t *testing.T) {
	redactable := model.Redactable[string]{
		Value:     "secret-data",
		VisibleTo: map[model.PlayerID]bool{},
	}

	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}

func TestRedact_NonRedactableValue(t *testing.T) {
	data := map[string]interface{}{
		"name":  "test",
		"value": 123,
	}

	result, err := Redact(data, "player1")
	require.NoError(t, err)

	var decoded map[string]interface{}
	err = json.Unmarshal(result, &decoded)
	require.NoError(t, err)

	assert.Equal(t, "test", decoded["name"])
	assert.Equal(t, float64(123), decoded["value"])
}

func TestRedact_PlayerIDSubstring(t *testing.T) {
	redactable := model.Redactable[string]{
		Value: "secret-data",
		VisibleTo: map[model.PlayerID]bool{
			"player1":   true,
			"player11":  false,
			"xplayer1x": false,
		},
	}

	// Test that "player1" doesn't match "player11" or "xplayer1x"
	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"secret-data"`, string(result))

	result, err = Redact(redactable, "player11")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))

	result, err = Redact(redactable, "xplayer1x")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}
