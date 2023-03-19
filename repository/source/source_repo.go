package source

import (
	"context"
	"github.com/rudderstack_source_app/constants"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type SourceRepository struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewSourceRepository(mongoClient *mongo.Client) *SourceRepository {
	collectionObj := mongoClient.Database(constants.DATABASE).Collection(constants.SOURCE_COLLECTION)
	return &SourceRepository{
		mongoClient: mongoClient,
		collection:  collectionObj,
	}
}

func (r *SourceRepository) CreateSource(ctx context.Context, source *entity.Source) error {
	daoSource := util.SourceEntityToDao(source)
	daoSource.DateCreated = time.Now()
	daoSource.DateUpdated = time.Now()

	_, err := r.collection.InsertOne(ctx, daoSource)
	if err != nil {
		return err
	}
	return nil
}

func (r *SourceRepository) GetAllSourcesByType(ctx context.Context, sourceType string) (*[]entity.Source, error) {
	filter := bson.M{"type": sourceType}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sources []util.Source
	for cursor.Next(ctx) {
		var source util.Source
		err := cursor.Decode(&source)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	var sourceEntities []entity.Source

	for _, source := range sources {
		sourceEntity := util.SourceDaoToEntity(source)
		sourceEntities = append(sourceEntities, *sourceEntity)
	}

	return &sourceEntities, nil
}
