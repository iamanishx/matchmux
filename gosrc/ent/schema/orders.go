package schema

import "entgo.io/ent"

// Orders holds the schema definition for the Orders entity.
type Orders struct {
	ent.Schema
}

// Fields of the Orders.
func (Orders) Fields() []ent.Field {
	return nil
}

// Edges of the Orders.
func (Orders) Edges() []ent.Edge {
	return nil
}
