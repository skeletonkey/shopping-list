package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instanceGen.NewApp("api", "app")
	app.SetupApp(
		app.WithDependencies(
			"github.com/pressly/goose/v3",
			"github.com/labstack/echo/v4",
			"github.com/mattn/go-sqlite3",
		),
		app.WithGithubWorkflows("linter", "test"),
		app.WithMakefile("db"),
		app.WithGoVersion("1.24.4"),
		app.WithPackages(
			"db",
			"eventing",
			"server",
		),
	).Generate()

}
