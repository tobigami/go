package main

import (
	"go-api/internal/routes"
)

func main() {
	r := routes.NewRouter()
	r.Run(":1411")
}
