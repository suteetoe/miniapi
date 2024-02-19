package microservice

type IContext interface {
	Param(name string) string
	QueryParam(name string) string
	ReadInput() string

	Response(responseCode int, responseData interface{})
	ResponseError(responseCode int, errorMessage string)
}
