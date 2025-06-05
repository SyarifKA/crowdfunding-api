package main

import (
	"fmt"
	"os"
	"time"

	"github.com/SyarifKA/crowdfunding-api/pkg/env"
	"github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/SyarifKA/crowdfunding-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := env.Init()
	if err != nil {
		log.Fatal(err)
	}

	// initialize config log
	err = os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	logTimestamp := time.Now().Format("2006-01-02_15-04")
	logFile := fmt.Sprintf("logs/%s.log", logTimestamp)

	err = log.SetConfig(&log.Config{
		Formatter: &log.TextFormatter,
		Level:     log.TraceLevel,
		LogName:   logFile,
	})
	if err != nil {
		log.Fatal(err)
	}
	// db := lib.DB()
	// migrations.Run(db)
	r := gin.Default()
	routers.RoutersCombine(r)
	r.Run("localhost:8888")
}
