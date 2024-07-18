package imagine

import (
	"net/http"
)

type ImagineHandlerInterface interface {
	Imagine(w http.ResponseWriter, r *http.Request)
}

type ImagineHandler struct{}

func NewHandler() ImagineHandlerInterface {
	return &ImagineHandler{}
}

func (h *ImagineHandler) Imagine(w http.ResponseWriter, r *http.Request) {
	//
}
