package migrations

import (
	"time"

	"github.com/SyarifKA/crowdfunding-api/models"
	log "github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
	)

	user := []models.User{{
		ID:        uuid.New().String(),
		Name:      "John Doe",
		Email:     "john12@example.com",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
	}, {
		ID:        uuid.New().String(),
		Name:      "syarif",
		Email:     "john10@example.com",
		Password:  "hashed_password",
		CreatedAt: time.Now(),
	},
	}

	// Insert ke database
	if err := db.Create(&user).Error; err != nil {
		panic("gagal insert user: " + err.Error())
	}

	log.Debug("Automigrate sukses!")
}
