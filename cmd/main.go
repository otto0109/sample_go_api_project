package main

import (
	"event_service/api/routes"
)

func main() {
	router := routes.NewRouter()
	router.Run()
}
