package main

import (
	"awesomeProject1/internal/list/service"
	router "awesomeProject1/pkg/app/router"
	"awesomeProject1/pkg/config"
	lapp "awesomeProject1/pkg/list"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	lapp.App = lapp.New()
	router.Load(lapp.App.Router)
	service.Init()
	lapp.App.Run()
}
