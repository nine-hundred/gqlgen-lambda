package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			Annotations(
				entgql.OrderField("TEXT"),
			).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Annotations(
				entgql.OrderField("STATUS"),
			).
			Default("IN_PROGRESS"),
		field.Int("priority").
			Annotations(
				entgql.OrderField("PRIORITY"),
			).
			Default(0),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Todo.Type).
			Unique().
			From("children"),
	}
}
