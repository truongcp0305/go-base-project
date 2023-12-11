package registry

import (
	"go-project/adapters/controller"

	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type registry struct {
	db *mongo.Collection
	es *elasticsearch.TypedClient
}

type Register interface {
	NewAppController() controller.AppController
}

func New(db *mongo.Collection, es *elasticsearch.TypedClient) Register {
	return &registry{
		db: db,
		es: es,
	}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
		Cmd:  r.NewCommandController(),
	}
}
