package sourcetemplate

import (
	"context"
	"github.com/rudderstack_source_app/entity"
)

type SourceTemplateServiceInterface interface {
	AddSourceTemplate(ctx context.Context, template entity.SourceTemplate, userId int64) (string, error)
	GetSourceTemplateByType(ctx context.Context, templateType string) (*entity.SourceTemplate, error)
}
