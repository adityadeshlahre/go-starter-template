package repository

import "github.com/adityadeshlahre/go-starter-template/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
