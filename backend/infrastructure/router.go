package infrastructure

import (
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/controllers"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/infrastructure/database"
	"github.com/labstack/echo"
	"net/http"
)

var Router *echo.Echo

func init() {
	router := echo.New()

	database.Connect()
	db := database.GetDB()
	userController := controllers.NewUserController(db)

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!")
	})
	router.GET("/users", func(c echo.Context) error {
		return userController.GetAllUsers(c)
	})
	router.GET("/createUser", func(c echo.Context) error {
		return userController.CreateUser(c)
	})
	Router = router
}
