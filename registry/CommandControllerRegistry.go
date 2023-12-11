package registry

import (
	"go-project/adapters/controller"
	"go-project/usecase/service"
)

func (r *registry) NewCommandController() controller.CommandController {
	return controller.NewCommandController(
		*service.NewCommandService(),
	)
}
