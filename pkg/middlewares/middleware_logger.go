package middlewares

import (
	"miniapi/pkg/config"
	"miniapi/pkg/logger"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type MiddlewareMetricsCb func(err error)

type IMiddlewareManager interface {
	RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type MiddlewareManager struct {
	log       logger.ILogger
	cfg       config.IConfig
	metricsCb MiddlewareMetricsCb
}

func NewMiddlewareManager(log logger.ILogger, cfg config.IConfig, metricsCb MiddlewareMetricsCb) IMiddlewareManager {
	return &MiddlewareManager{log: log, cfg: cfg, metricsCb: metricsCb}
}

func (mw *MiddlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		start := time.Now()
		err := next(ctx)

		req := ctx.Request()
		res := ctx.Response()
		status := res.Status
		size := res.Size
		s := time.Since(start)

		ip := ctx.RealIP()
		user := ""

		// userInfo := ctx.Get("UserInfo")
		// if userInfo != nil {
		// 	user = userInfo.(microserviceModel.UserInfo).Username
		// } else {
		// 	tempUserName := ctx.Get("username")
		// 	if tempUserName != nil {
		// 		user = tempUserName.(string)
		// 	}
		// }

		httpErr, isHttpError := err.(*echo.HTTPError)

		if isHttpError {
			status = httpErr.Code
		}

		if !mw.CheckIgnoredURI(ctx.Request().RequestURI, mw.cfg.HttpConfig().IgnoreLogUrls()) {
			mw.log.HttpMiddlewareAccessLogger(req.Method, req.URL.String(), status, size, s, ip, user)
		}

		mw.metricsCb(err)
		return err
	}
}

func (mw *MiddlewareManager) CheckIgnoredURI(requestURI string, uriList []string) bool {
	for _, s := range uriList {
		if strings.Contains(requestURI, s) {
			return true
		}
	}
	return false
}
