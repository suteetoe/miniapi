package main

import (
	"miniapi/internal/demoapi"
	"miniapi/pkg/config"
	"miniapi/pkg/microservice"
)

func main() {
	// Call the function

	cfg := config.NewConfig()
	ms := microservice.NewMicroservice(cfg)

	ms.RegisterHttp(demoapi.NewDemoApi(ms, cfg))
	ms.Start()
}
