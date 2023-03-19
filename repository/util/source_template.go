package util

import (
	"github.com/rudderstack_source_app/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SourceTemplate struct {
	ID     primitive.ObjectID     `bson:"_id,omitempty"`
	Type   string                 `bson:"type"`
	Fields map[string]interface{} `bson:"fields"`
}

func SourceTemplateDaoToEntity(sourceTemplate SourceTemplate) *entity.SourceTemplate {
	return &entity.SourceTemplate{
		Type:   sourceTemplate.Type,
		Fields: sourceTemplate.Fields,
	}
}

func SourceTemplateEntityToDao(e *entity.SourceTemplate) *SourceTemplate {
	return &SourceTemplate{
		Type:   e.Type,
		Fields: e.Fields,
	}
}
