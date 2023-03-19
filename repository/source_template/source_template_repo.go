package source_template

import (
	"context"
	"github.com/rudderstack_source_app/constants"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SourceTemplateRepository struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewSourceTemplateRepository(mongoClient *mongo.Client) *SourceTemplateRepository {
	collectionObj := mongoClient.Database(constants.DATABASE).Collection(constants.SOURCE_TEMPLATE_COLLECTION)
	return &SourceTemplateRepository{
		mongoClient: mongoClient,
		collection:  collectionObj,
	}
}

func (r *SourceTemplateRepository) GetSourceTemplateByType(ctx context.Context, templateType string) (*entity.SourceTemplate, error) {
	filter := bson.M{"type": templateType}
	var sourceTemplate util.SourceTemplate
	err := r.collection.FindOne(ctx, filter).Decode(&sourceTemplate)
	if err != nil {
		return nil, err
	}
	return util.SourceTemplateDaoToEntity(sourceTemplate), nil
}

func (r *SourceTemplateRepository) InsertSourceTemplate(ctx context.Context, template *entity.SourceTemplate) (string, error) {

	// Insert the new source template document
	daoTemplate := util.SourceTemplateEntityToDao(template)

	result, err := r.collection.InsertOne(ctx, daoTemplate)
	if err != nil {
		return "", err
	}

	// Extract the inserted document ID
	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *SourceTemplateRepository) GetAllSourceTemplates(ctx context.Context) (*[]string, error) {
	var types []string
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var template util.SourceTemplate
		if err := cursor.Decode(&template); err != nil {
			return nil, err
		}
		types = append(types, template.Type)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &types, nil
}
