package microservice

import (
	"errors"
	"fmt"
	"miniapi/pkg/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type IMicroserviceHTTP interface {
	RegisterHttp()
}

func httpPanicHandler(ms *Microservice, c echo.Context) {
	r := recover()
	if r != nil {
		ms.Logger.HttpInternalServerErrorLogger(c.Path(), errors.New(fmt.Sprintf("%v", r)), time.Now())
		// fmt.Println("Panic: ", r)

		c.JSON(http.StatusInternalServerError, models.RestErrorResponse{
			ErrStatus:  500,
			ErrMessage: "Internal Server Error",
			Timestamp:  time.Now(),
		})
	}
}

// GET register service endpoint for HTTP GET
func (ms *Microservice) GET(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {

	ms.Logger.Debugf("Register HTTP Handler GET \"%s\".", path)
	ms.echo.GET(path, func(c echo.Context) error {
		defer httpPanicHandler(ms, c)
		return h(NewHTTPContext(ms, c))
	}, m...)
}

// POST register service endpoint for HTTP POST
func (ms *Microservice) POST(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {

	fullPath := ms.config.HttpConfig().PathPrefix() + path
	ms.Logger.Debugf("Register HTTP Handler POST \"%s\".", fullPath)
	ms.echo.POST(fullPath, func(c echo.Context) error {
		defer httpPanicHandler(ms, c)
		return h(NewHTTPContext(ms, c))
	}, m...)
}

// PUT register service endpoint for HTTP PUT
func (ms *Microservice) PUT(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	fullPath := ms.config.HttpConfig().PathPrefix() + path
	ms.Logger.Debugf("Register HTTP Handler PUT \"%s\".", fullPath)
	ms.echo.PUT(fullPath, func(c echo.Context) error {
		defer httpPanicHandler(ms, c)
		return h(NewHTTPContext(ms, c))
	}, m...)
}

// PATCH register service endpoint for HTTP PATCH
func (ms *Microservice) PATCH(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	fullPath := ms.config.HttpConfig().PathPrefix() + path
	ms.Logger.Debugf("Register HTTP Handler PATCH \"%s\".", fullPath)
	ms.echo.PATCH(fullPath, func(c echo.Context) error {
		defer httpPanicHandler(ms, c)
		return h(NewHTTPContext(ms, c))
	}, m...)
}

// DELETE register service endpoint for HTTP DELETE
func (ms *Microservice) DELETE(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc) {
	fullPath := ms.config.HttpConfig().PathPrefix() + path
	ms.Logger.Debugf("Register HTTP Handler DELETE \"%s\".", fullPath)
	ms.echo.DELETE(fullPath, func(c echo.Context) error {
		defer httpPanicHandler(ms, c)
		return h(NewHTTPContext(ms, c))
	}, m...)
}
