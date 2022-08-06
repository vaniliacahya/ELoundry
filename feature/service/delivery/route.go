package delivery

import (
	"RESTAPILoundry/config"
	"RESTAPILoundry/domain"
	"RESTAPILoundry/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteServices(e *echo.Echo, sh domain.ServiceHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/loundry", sh.InsertServ(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/loundry/:id", sh.UpdateServ(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/loundry/:id", sh.DeleteServ(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/loundry", sh.GetAllServ())
	e.GET("/loundry/:id", sh.GetServID())
}
