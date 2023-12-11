package outgoing

import "go-project/model"

type LoginOutgoing struct {
	Data string `json:"data" example:"ok"`
}

type ModelInternalErr struct {
	Error string `json:"error" example:"Internal sever error"`
}

type ModelBadRequestErr struct {
	Error string `json:"error" example:"Invalid query params"`
}

type RegisterOutgoing struct {
	Data model.User `json:"data"`
}
