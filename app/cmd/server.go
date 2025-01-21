package cmd

import (
	"context"

	"github.com/golang-documented-todo-api/app/api"
	"github.com/golang-documented-todo-api/app/datasources"
)

func RunServer(ctx context.Context, datasource *datasources.DataSources) error {
	app := api.CreateApp(ctx, datasource)
	return app.Listen(":3000")
}
