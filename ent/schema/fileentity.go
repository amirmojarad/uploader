package schema

import "entgo.io/ent"

// FileEntity holds the schema definition for the FileEntity entity.
type FileEntity struct {
	ent.Schema
}

// Fields of the FileEntity.
func (FileEntity) Fields() []ent.Field {
	return nil
}

// Edges of the FileEntity.
func (FileEntity) Edges() []ent.Edge {
	return nil
}
