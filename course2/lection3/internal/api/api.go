package api

import "fmt"

type API struct {
	config *Config
}

func New() *API {
	return &API{}
}

func (s *API) Start() error {
	fmt.Println("Server is running")
	return nil
}
