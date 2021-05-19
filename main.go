package main

import (
	"StarWarsBackEnd/routes"
)

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
