package main

import (
	"fmt"
	"github.com/dengzii/blog/bootstrap"
	"github.com/dengzii/blog/conf"
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models"
	"github.com/dengzii/blog/routers"
)

var app *bootstrap.Bootstrapper

func main() {

	db.Init()
	models.Init()
	app = bootstrap.New(
		conf.Get().Iris.AppName,
		conf.Get().Iris.Owner,
		true)

	app.SetupViews("./views")
	app.Bootstrap()

	routers.Setup(app)
	app.Listen(fmt.Sprintf(":%d", conf.Get().Iris.Port))
}
