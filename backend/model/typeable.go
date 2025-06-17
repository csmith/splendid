package model

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
