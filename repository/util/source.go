package util

import (
	"github.com/rudderstack_source_app/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Source struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	UserId      int64                  `bson:"user_id,omitempty"`
	Type        string                 `bson:"type"`
	Data        map[string]interface{} `bson:"data"`
	DateCreated time.Time              `bson:"date_created,omitempty"`
	DateUpdated time.Time              `bson:"date_updated,omitempty"`
}

func SourceDaoToEntity(source Source) *entity.Source {
	return &entity.Source{
		UserId:      source.UserId,
		Type:        source.Type,
		Data:        source.Data,
		DateCreated: source.DateCreated,
		DateUpdated: source.DateUpdated,
	}
}

func SourceEntityToDao(source *entity.Source) *Source {
	return &Source{
		UserId:      source.UserId,
		Type:        source.Type,
		Data:        source.Data,
		DateCreated: source.DateCreated,
		DateUpdated: source.DateUpdated,
	}
}
