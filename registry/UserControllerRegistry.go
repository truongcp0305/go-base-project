package registry

import (
	"go-project/adapters/controller"
	"go-project/adapters/repository"
	"go-project/usecase/service"
)

func (r *registry) NewUserController() controller.UserController {
	db := repository.NewUserRepository(r.db, r.es)
	return controller.NewUserController(
		*service.NewUserService(db),
	)
}
