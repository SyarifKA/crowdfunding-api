package main

import (
	"os"

	"github.com/SyarifKA/crowdfunding-api/lib"
	"github.com/SyarifKA/crowdfunding-api/migrations"
	"github.com/SyarifKA/crowdfunding-api/pkg/env"
	"github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/SyarifKA/crowdfunding-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := env.Init()
	if err != nil {
		log.Fatal(err)
	}

	// initialize config log
	err = os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// logTimestamp := time.Now().Format("2006-01-02_15-04")
	// logFile := fmt.Sprintf("logs/%s.log", logTimestamp)

	// err = log.SetConfig(&log.Config{
	// 	Formatter: &log.TextFormatter,
	// 	Level:     log.TraceLevel,
	// 	LogName:   logFile,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Init logger config
	err = log.InitLogger(&log.Config{
		Formatter: &log.TextFormatter,
		Level:     log.InfoLevel,
		LogName:   "application.log",
	})
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("RUN_MIGRATION") == "true" {
		db := lib.DB()
		migrations.Run(db)
	}

	// r.Use(middlewares.LogMiddlware)
	routers.RoutersCombine(r)
	r.Run("localhost:8888")
}
