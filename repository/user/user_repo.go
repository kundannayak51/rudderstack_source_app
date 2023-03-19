package user

import (
	"context"
	"github.com/rudderstack_source_app/constants"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/repository/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepository struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewUserRepository(mongoClient *mongo.Client) *UserRepository {
	collectionObj := mongoClient.Database(constants.DATABASE).Collection(constants.USER_COLLECTION)
	return &UserRepository{
		mongoClient: mongoClient,
		collection:  collectionObj,
	}
}

func (r *UserRepository) GetUserByUserId(ctx context.Context, userId int64) (*entity.User, error) {
	filter := bson.M{"user_id": userId}
	var user util.User
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return util.UserDaoToEntity(user), nil
}

func (r *UserRepository) CreateOrUpdateUser(ctx context.Context, user *entity.User) error {
	filter := bson.M{
		"user_id": user.UserId,
	}

	data := util.UserEntityToDao(user)
	data.UpdatedAt = time.Now()
	update := bson.M{
		"$set": data,
	}
	upsert := true
	_, err := r.collection.UpdateOne(ctx, filter, update, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}
	return nil
}
