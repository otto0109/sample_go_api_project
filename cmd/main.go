package main

import (
	"sample_go_api_project/api/routes"
)

func main() {
	router := routes.NewRouter()
	router.Run()
}
