package microservice

import (
	"context"
	"miniapi/pkg/config"
	"miniapi/pkg/logger"
	"miniapi/pkg/middlewares"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type IMicroservice interface {
	Start() error
	Cleanup() error
	RegisterHttp(http IMicroserviceHTTP)

	GET(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	POST(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	PUT(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	PATCH(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
	DELETE(path string, h ServiceHandleFunc, m ...echo.MiddlewareFunc)
}

type Microservice struct {
	config            config.IConfig
	Logger            logger.ILogger
	exitChannel       chan bool
	echo              *echo.Echo
	middlewareManager middlewares.IMiddlewareManager
}

type ServiceHandleFunc func(context IContext) error

func NewMicroservice(cfg config.IConfig) IMicroservice {

	loggerConfig := config.NewLoggerConfig()
	logger := logger.NewAppLogger(loggerConfig)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	m := &Microservice{
		echo:   e,
		config: cfg,
		Logger: logger,
	}

	m.middlewareManager = middlewares.NewMiddlewareManager(logger, cfg, m.getHttpMetricsCb())

	return m
}

func (ms *Microservice) Start() error {

	ms.Logger.Debugf("Start App: %s Mode: %s", ms.config.ApplicationName(), ms.config.Mode())

	// if ms.Mode == "development" {
	// 	// register swagger api spec
	// 	ms.echo.Static("/swagger/doc.json", "./../../api/swagger/swagger.json")
	// }
	httpN := len(ms.echo.Routes())
	var exitHTTP chan bool
	if httpN > 0 {
		exitHTTP = make(chan bool, 1)
		go func() {
			ms.startHTTP(exitHTTP)
		}()

	}

	osQuit := make(chan os.Signal, 1)
	ms.exitChannel = make(chan bool, 1)
	signal.Notify(osQuit, syscall.SIGTERM, syscall.SIGINT)
	exit := false
	for {
		if exit {
			break
		}
		select {
		case <-osQuit:
			// Exit from HTTP as well
			if exitHTTP != nil {
				exitHTTP <- true
			}
			exit = true
		case <-ms.exitChannel:
			// Exit from HTTP as well
			if exitHTTP != nil {
				exitHTTP <- true
			}
			exit = true
		}
	}

	defer ms.Cleanup()

	return nil
}

func (ms *Microservice) Cleanup() error {

	ms.Logger.Info("Stop Service Cleanup System.")
	return nil
}

func (ms *Microservice) RegisterHttp(http IMicroserviceHTTP) {
	http.RegisterHttp()
}

func (ms *Microservice) startHTTP(exitChannel chan bool) error {

	if ms.middlewareManager != nil {
		ms.echo.Use(ms.middlewareManager.RequestLoggerMiddleware)
	}

	port := ms.config.HttpConfig().Port()

	// Caller can exit by sending value to exitChannel
	go func() {
		<-exitChannel
		ms.stopHTTP()
	}()

	ms.Logger.Infof("Listening: %v Entrypoint: %v ", port, ms.config.HttpConfig().PathPrefix())

	err := ms.echo.Start("0.0.0.0:" + port)

	if err == nil {
		ms.Logger.Error("Failed After Start", err)
	}

	return err
}

func (ms *Microservice) stopHTTP() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ms.echo.Shutdown(ctx)
}
