package delivery

import (
	"wa/config"
	"wa/domain"
	"wa/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteMessage(e *echo.Echo, mh domain.MessageHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/message", mh.InsertMessage(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/message/:id", mh.DeleteMessage(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}
