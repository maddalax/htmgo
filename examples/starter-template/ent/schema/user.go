package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("occupation").Optional(),
		field.String("email").Optional(),
		field.String("password").Sensitive().Optional(),
		field.String("test").Optional(),
		field.String("test2").Optional(),
		field.String("test4").Optional(),
		field.Time("created_at").
			Default(time.Now).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
		field.Time("updated_at").
			Default(time.Now).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
