package source

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type SourceRepo interface {
	CreateSource(ctx context.Context, source *entity.Source) error
	GetAllSourcesByType(ctx context.Context, sourceType string) (*[]entity.Source, error)
}
