package repository

import (
	"github.com/SyarifKA/crowdfunding-api/entity"
	"github.com/SyarifKA/crowdfunding-api/lib"
	log "github.com/SyarifKA/crowdfunding-api/pkg/log"
)

func FindAllUsers() ([]entity.User, error) {
	db := lib.DB()
	var users []entity.User

	result := db.Select("name", "email").Find(&users)
	if result.Error != nil {
		log.Error("Error get all users:", result.Error)
		return nil, result.Error
	}

	log.Debug("get all users repository success")
	return users, nil
}
