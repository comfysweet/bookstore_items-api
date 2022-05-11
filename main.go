package main

import (
	"github.com/comfysweet/bookstore_items-api/app"
	"os"
)

func main() {
	os.Setenv("LOG_LEVEL", "info")
	app.StartApplication()
}
