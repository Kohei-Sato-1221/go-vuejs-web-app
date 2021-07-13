package infrastructure

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/controllers"
	"github.com/labstack/echo"
	"net/http"
)

var Router *echo.Echo

func init() {
	router := echo.New()

	userController := controllers.NewUserController()

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})
	router.GET("/users", func(c echo.Context) error {
		data := userController.GetAllUsers()
		return c.String(http.StatusOK, data)
	})
	Router = router
}