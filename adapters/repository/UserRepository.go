package repository

import (
	"context"
	"encoding/json"
	"go-project/model"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

func (r *UserRepository) FindUserByUserName(userName string) ([]model.User, error) {
	result := []model.User{}
	res, err := r.es.Search().Index(os.Getenv("INDEX")).Request(
		&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"userName": {Query: userName},
				},
			},
		},
	).Do(context.Background())
	if err != nil {
		return result, err
	}
	for _, i := range res.Hits.Hits {
		u := model.User{}
		err := json.Unmarshal(i.Source_, &u)
		if err != nil {
			return result, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.db.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepository) FindUser(userName string, password string) ([]model.User, error) {
	result := []model.User{}
	res, err := r.es.Search().Index(os.Getenv("INDEX")).Request(
		&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"userName": {Query: userName},
					"password": {Query: password},
				},
			},
		},
	).Do(context.Background())
	if err != nil {
		return result, err
	}
	for _, i := range res.Hits.Hits {
		u := model.User{}
		err := json.Unmarshal(i.Source_, &u)
		if err != nil {
			return result, err
		}
		result = append(result, u)
	}
	return result, nil
}
