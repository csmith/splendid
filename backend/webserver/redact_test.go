package webserver

import (
	"encoding/json"
	"testing"

	"github.com/csmith/splendid/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRedact_BasicVisiblePlayer(t *testing.T) {
	redactable := model.NewRedactable("secret-data", "player1")

	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"secret-data"`, string(result))
}

func TestRedact_CardVisiblePlayer(t *testing.T) {
	redactable := model.NewRedactable("Princess", "player1")

	result, err := Redact(redactable, "player1")
	require.NoError(t, err)
	assert.Equal(t, `"Princess"`, string(result))
}

func TestRedact_BasicNonVisiblePlayer(t *testing.T) {
	redactable := model.NewRedactable("secret-data", "player1")

	result, err := Redact(redactable, "player2")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}

func TestRedact_UnknownPlayer(t *testing.T) {
	redactable := model.NewRedactable("secret-data", "player1")

	result, err := Redact(redactable, "player3")
	require.NoError(t, err)
	assert.Equal(t, `"REDACTED"`, string(result))
}

func TestRedact_NestedRedactable(t *testing.T) {
	innerRedactable := model.NewRedactable("inner-secret", "player1")
	outerRedactable := model.NewRedactable(innerRedactable, "player1", "player2")

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
		Public:  "everyone-can-see",
		Secret1: model.NewRedactable("secret-string", "player1"),
		Secret2: model.NewRedactable(42, "player2"),
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
	redactable := model.NewRedactable[string]("secret-data")

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
	redactable := model.NewRedactable("secret-data", "player1")

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
