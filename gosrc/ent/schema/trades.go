package schema

import "entgo.io/ent"

// Trades holds the schema definition for the Trades entity.
type Trades struct {
	ent.Schema
}

// Fields of the Trades.
func (Trades) Fields() []ent.Field {
	return nil
}

// Edges of the Trades.
func (Trades) Edges() []ent.Edge {
	return nil
}
