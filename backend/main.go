package main

import (
	"fmt"
	"github.com/Kohei-Sato-1221/go-vuejs-web-app/backend/infrastructure"
)

func main() {
	fmt.Println("Started!")
	infrastructure.Router.Start(":8080")
}
