package lib

import (
	log "github.com/SyarifKA/crowdfunding-api/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "host=localhost password=1 user=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	log.Debug("db connect sukses!")
	return db
}
