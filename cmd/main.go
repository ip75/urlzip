package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ip75/urlzip/pkg/config"
	"github.com/ip75/urlzip/pkg/core"
	"github.com/ip75/urlzip/pkg/repository"
	"github.com/ip75/urlzip/pkg/rest"
)

func main() {

	l := log.Default()

	cfg := config.ReadConfig()

	db, err := repository.Connect(cfg)
	if err != nil {
		l.Fatalln("connect to database failed: ", err)
		return
	}

	err = db.Initialize()
	if err != nil {
		l.Fatalln("initialize database failed: ", err)
		return
	}

	r := gin.Default()
	rest.RaiseHanders(r, core.NewUrlZipCore(db))

	if err := r.Run(":" + cfg.Listen); err != nil {
		l.Fatalln("service failed:", err)
		return
	}
	l.Println("service stoped...")
}
