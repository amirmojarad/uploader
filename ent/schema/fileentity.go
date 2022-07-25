package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FileEntity holds the schema definition for the FileEntity entity.
type FileEntity struct {
	ent.Schema
}

// Fields of the FileEntity.
func (FileEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("type"),
		field.Int64("size"),
		field.String("url"),
	}
}

// Edges of the FileEntity.
func (FileEntity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("files").
			Unique(),
	}
}
