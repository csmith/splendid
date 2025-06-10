package serialization

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Specifier uniquely identifies a type with domain, category, and name
type Specifier struct {
	Domain   string // Game/domain identifier (e.g., "dyhat", "poker")
	Category string // Type category (e.g., "a" for actions, "e" for events)
	Name     string // Type name (e.g., "add_player", "card_dealt")
}

// String returns the full qualified type name
func (s Specifier) String() string {
	return s.Domain + ":" + s.Category + ":" + s.Name
}

// Typeable represents any type that has a Type() method returning a Specifier
type Typeable interface {
	Type() Specifier
}

// Box is a generic wrapper that handles JSON marshalling/unmarshalling for types with Type() methods
type Box[T Typeable] struct {
	Value T
}

// typeRegistry maps "domain:category:name" keys to constructor functions
var typeRegistry = make(map[string]func() Typeable)

// RegisterType registers a type for automatic unmarshalling
// The type's Type() method provides the domain, category, and name
func RegisterType[T Typeable](example T) {
	spec := example.Type()
	key := spec.String()

	typeRegistry[key] = func() Typeable {
		// Create a new zero value of the same type as the example
		typ := reflect.TypeOf(example)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		return reflect.New(typ).Interface().(Typeable)
	}
}

// MarshalJSON implements json.Marshaler
func (b *Box[T]) MarshalJSON() ([]byte, error) {
	// Get the full type specifier from the value
	spec := b.Value.Type()

	// First marshal the value itself
	valueBytes, err := json.Marshal(b.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal value: %w", err)
	}

	// Then unmarshal it into a map to add the type field
	var valueMap map[string]interface{}
	if err := json.Unmarshal(valueBytes, &valueMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal value as map: %w", err)
	}

	// Add the type field
	valueMap["type"] = spec.String()

	// Marshal the final result
	return json.Marshal(valueMap)
}

// UnmarshalJSON implements json.Unmarshaler
func (b *Box[T]) UnmarshalJSON(data []byte) error {
	// First, unmarshal just the type field
	var typeInfo struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &typeInfo); err != nil {
		return fmt.Errorf("failed to unmarshal type field: %w", err)
	}

	// Look up the constructor in the registry
	constructor, exists := typeRegistry[typeInfo.Type]
	if !exists {
		return fmt.Errorf("unknown type: %s", typeInfo.Type)
	}

	// Create a new instance of the correct type
	value := constructor()

	// Unmarshal the full JSON into the concrete type
	if err := json.Unmarshal(data, value); err != nil {
		return fmt.Errorf("failed to unmarshal %s: %w", typeInfo.Type, err)
	}

	// Type assert to T and store in the box
	typedValue, ok := value.(T)
	if !ok {
		return fmt.Errorf("type assertion failed: expected %T, got %T", *new(T), value)
	}

	b.Value = typedValue
	return nil
}

func (b *Box[T]) String() string {
	return fmt.Sprintf("%v", b.Value)
}

// Unmarshal is a convenience function for unmarshalling into a Box
func Unmarshal[T Typeable](data []byte) (*Box[T], error) {
	var box Box[T]
	err := box.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}
	return &box, nil
}
