package handler

import (
	"github.com/adityadeshlahre/go-starter-template/internal/server"
	"github.com/adityadeshlahre/go-starter-template/internal/service"
)

type Handlers struct {
	Health  *HealthHandler
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health:  NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
