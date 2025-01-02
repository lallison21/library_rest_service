package main

import "github.com/lallison21/library_rest_service/internal/application"

func main() {
	app := application.New()

	app.Run()
}
