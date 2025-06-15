package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Otp holds the schema definition for the Otp entity.
type Otp struct {
	ent.Schema
}

// Fields of the Otp.
func (Otp) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").NotEmpty(),
		field.Time("expires_at"),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the Otp.
func (Otp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", Users.Type).
			Ref("otp").
			Unique().
			Field("user_id").
			Required(),
	}
}
