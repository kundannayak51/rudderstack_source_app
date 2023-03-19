package source_template

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type SourceTemplateRepo interface {
	GetSourceTemplateByType(ctx context.Context, templateType string) (*entity.SourceTemplate, error)
	InsertSourceTemplate(ctx context.Context, template *entity.SourceTemplate) (string, error)
}
