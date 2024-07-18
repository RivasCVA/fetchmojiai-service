package server

import (
	"net/http"

	"github.com/RivasCVA/fetchmojiai-service/api"
	"github.com/RivasCVA/fetchmojiai-service/server/handler/imagine"
)

// Server implements the generated "ServerInterface" from OpenAPI.
// It forwards all calls to their respective handlers.
type Server struct {
	imagineHandler imagine.ImagineHandlerInterface
}

func NewServer(imagineHandler imagine.ImagineHandlerInterface) api.ServerInterface {
	return &Server{
		imagineHandler: imagineHandler,
	}
}

func (s *Server) Imagine(w http.ResponseWriter, r *http.Request) {
	s.imagineHandler.Imagine(w, r)
}
