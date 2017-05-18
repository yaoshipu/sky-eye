package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"skyeye.paypal.com/common/config"
	"skyeye.paypal.com/service"
	"skyeye.paypal.com/skyeye/router"
)

const (
	MiddlewareLogFormat = `{"time": ${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"method":"${method}","uri":"${uri}","status":${status}}` + "\n"

	LogFormat = `${time_rfc3339}" [${level}] ${long_file}:${line}`
)

func main() {

	config.Init("f", "skyeye", "skyeye.conf")

	var cfg *service.ServiceConfig
	err := config.Load(&cfg)
	if err != nil {
		log.Fatal("load SkyEye config error:", err)
		return
	}

	service.Init(cfg)

	log.SetHeader(LogFormat)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: MiddlewareLogFormat,
	}))
	e.Use(middleware.RequestID())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())

	router.AddRouters(e)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))

}
