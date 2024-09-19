package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ent.Schema
}

// Fields of the User.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").
			Default("unknown"),
		field.Time("created_at").Default(time.Now).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
		field.Time("updated_at").Default(time.Now).Annotations(
			entsql.Default("CURRENT_TIMESTAMP"),
		),
		field.Time("completed_at").Optional().Nillable(),
		field.Strings("tags").Optional(),
	}
}

// Edges of the User.
func (Task) Edges() []ent.Edge {
	return nil
}
