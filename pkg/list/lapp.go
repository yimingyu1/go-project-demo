package lapp

import (
	"awesomeProject1/internal/list/service/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	MODE_DEBUG   string = "debug"
	MODE_RELEASE string = "release"
	MODE_TEST    string = "test"
)

var App *Appplication

type Appplication struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func New() *Appplication {
	app := new(Appplication)
	app.DB = model.Init()
	model.MigrateDB()
	gin.SetMode(gin.ReleaseMode)
	app.Router = gin.Default()
	return app
}

func (a *Appplication) Run() {
	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("app.addr"))
	srv := &http.Server{
		Addr:    viper.GetString("app.addr"),
		Handler: a.Router,
	}
	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s", err.Error())
	}
}
