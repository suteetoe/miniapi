package demoapi

import (
	"fmt"
	"miniapi/pkg/config"
	"miniapi/pkg/microservice"
	"miniapi/pkg/models"
)

type IDemoApi interface {
	RegisterHttp()
}

type DemoStruct struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type DemoApi struct {
	ms  microservice.IMicroservice
	cfg config.IConfig
}

var demos = []DemoStruct{{Code: "001", Name: "John", Age: 20}, {Code: "002", Name: "Doe", Age: 30}, {Code: "003", Name: "Smith", Age: 40}}

func NewDemoApi(ms microservice.IMicroservice, config config.IConfig) IDemoApi {
	return &DemoApi{
		ms:  ms,
		cfg: config,
	}
}

func (api *DemoApi) RegisterHttp() {
	// api.ms.GET().GET("/demo", api.demo)
	api.ms.GET("/demo", api.ListDemo)
	api.ms.GET("/demo/:code", api.GetDemo)
}

func (api *DemoApi) ListDemo(ctx microservice.IContext) error {

	// response list of models
	ctx.Response(200, models.ApiResponse{
		Success: true,
		Data:    demos,
	})
	return nil
}

func (api *DemoApi) GetDemo(ctx microservice.IContext) error {

	code := ctx.Param("code")

	getDemo := DemoStruct{}
	for _, d := range demos {
		if d.Code == code {
			getDemo = d
		}
	}

	if getDemo.Code == "" {
		errorMsg := fmt.Sprintf("Not Found %s", code)
		ctx.ResponseError(404, errorMsg)
		return nil
	}

	ctx.Response(200, models.ApiResponse{
		Success: true,
		Data:    getDemo,
	})
	return nil
}
