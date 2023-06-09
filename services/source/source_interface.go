package source

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type SourceServiceInterface interface {
	CreateSource(ctx context.Context, source entity.Source) error
	GetAllSources(ctx context.Context, sourceType string) (*[]entity.Source, error)
}
