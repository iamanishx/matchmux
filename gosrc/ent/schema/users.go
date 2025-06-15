package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
        field.String("password").Sensitive(), 
        field.String("email").Unique(),
		field.String("phone").Unique(),
        field.Time("created_at").Default(time.Now),
    }
}
// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
