package servd

import (
	"fmt"

	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/acme/autocert"
)

// StartServer starts server with given options like autoTLS and port
func StartServer(port int, autoTLS bool) error {
	srvAddress := fmt.Sprintf(":%s", strconv.Itoa(port))
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	addHandlers(e)
	var err error
	if autoTLS {
		e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
		err = e.StartAutoTLS(srvAddress)
	} else {
		err = e.Start(srvAddress)
	}
	e.Logger.Fatal(err)

	return err
}

func addHandlers(e *echo.Echo) {
	h := newHandler()
	e.GET("/", h.GetFileData)
	e.POST("/", h.AddFileData)
}
