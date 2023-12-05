package repository

import (
	"context"
	"fmt"
	"go-project/model"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
	es *elasticsearch.TypedClient
}

func NewUserRepository(db *mongo.Collection, es *elasticsearch.TypedClient) *UserRepository {
	return &UserRepository{
		db: db,
		es: es,
	}
}

func (r *UserRepository) FindUserWithUserName(userName string) ([]model.User, error) {
	filter := bson.D{{Key: "username", Value: userName}}
	cursor, err := r.db.Find(context.Background(), filter)
	if err != nil {
		return []model.User{}, err
	}
	defer cursor.Close(context.Background())
	r.find(userName)
	users := []model.User{}
	for cursor.Next(context.Background()) {
		user := model.User{}
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (r *UserRepository) find(userName string) {
	res, err := r.es.Search().Index(os.Getenv("INDEX")).Request(
		&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"username": {Query: userName},
				},
			},
		},
	).Do(context.Background())
	fmt.Println(res, err)
	for _, i := range res.Hits.Hits {
		fmt.Println(i)
	}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.db.InsertOne(context.Background(), user)
	return err
}
