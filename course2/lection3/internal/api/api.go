package api

import "fmt"

type API struct {
	config *Config
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (s *API) Start() (err error) {
	fmt.Println("Server is running")
	return
}
