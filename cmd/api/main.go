package main

import (
	"github.com/phanorcoll/go_mongo/internal/api"
	"github.com/phanorcoll/go_mongo/pkg/config"
	"github.com/phanorcoll/go_mongo/pkg/data"
)

func main() {
	cfg := config.New()
	db := data.NewMongoConnection(cfg)
	defer db.Disconnect()
	application := api.New()
	application.Start()
}
