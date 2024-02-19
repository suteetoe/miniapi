package microservice

import (
	"io"

	"github.com/labstack/echo/v4"
)

type HTTPContext struct {
	ms *Microservice
	c  echo.Context
}

func NewHTTPContext(ms *Microservice, c echo.Context) *HTTPContext {
	return &HTTPContext{
		ms: ms,
		c:  c,
	}
}

func (ctx *HTTPContext) Param(name string) string {
	return ctx.c.Param(name)
}

func (ctx *HTTPContext) QueryParam(name string) string {
	return ctx.c.QueryParam(name)
}

func (ctx *HTTPContext) ReadInput() string {
	body, err := io.ReadAll(io.Reader(ctx.c.Request().Body))
	if err != nil {
		return ""
	}
	return string(body)
}

func (ctx *HTTPContext) Response(responseCode int, responseData interface{}) {
	ctx.c.JSON(responseCode, responseData)
}

func (ctx *HTTPContext) ResponseError(responseCode int, errorMessage string) {
	ctx.c.JSON(responseCode, map[string]interface{}{
		"success": false,
		"message": errorMessage,
	})
}
