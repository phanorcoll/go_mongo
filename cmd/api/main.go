package main

import "github.com/phanorcoll/go_mongo/internal/api"

func main() {
	application := api.New()
	application.Start()
}
