package main

import (
	"fmt"

	"github.com/fla-t/go-ing/UserService/internal/app"
)

func main() {
	// Use `true` for InMemory implementation, `false` for SQL implementation
	application := app.NewApp(true)

	fmt.Println("UserService running on http://localhost:8081")
	if err := application.Router.Run(":8081"); err != nil {
		panic(err)
	}
}
