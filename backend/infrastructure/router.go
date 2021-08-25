package infrastructure

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/adapter/controllers"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/graph"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/graph/generated"
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

	//For GraphQL
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	router.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	router.GET("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	router.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	fmt.Println("connect to http://localhost:8080/playground for GraphQL playground")
	fmt.Println("connect to http://localhost:8080/query for GraphQL Query")

	Router = router
}
