package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NodeTimeEntry holds the schema definition for the NodeTimeEntry entity.
type NodeTimeEntry struct {
	ent.Schema
}

// Fields of the NodeTimeEntry.
func (NodeTimeEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Int("nodeId"),
		field.Time("startTime"),
		field.Time("endTime"),
	}
}

// Edges of the NodeTimeEntry.
func (NodeTimeEntry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("time_entries").
			Unique(),
	}
}
