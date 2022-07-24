package crud

import (
	"context"
	"uploader/ent"
)

type Crud struct {
	Ctx    *context.Context
	Client *ent.Client
}
