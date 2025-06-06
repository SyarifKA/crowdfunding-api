package repository

import (
	"github.com/SyarifKA/crowdfunding-api/dtos"
	"github.com/SyarifKA/crowdfunding-api/entity"
	"github.com/SyarifKA/crowdfunding-api/lib"
	"github.com/SyarifKA/crowdfunding-api/models"
	log "github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/google/uuid"
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

func RegistUser(data dtos.RegistUser) (entity.RegistUser, error) {
	db := lib.DB()

	id := uuid.New().String()

	user := models.User{
		ID:       id,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password, // idealnya sudah di-hash sebelum ini
	}

	result := db.Select("id", "Name", "Email", "Password").Create(&user)
	if result.Error != nil {
		log.Error("Error create user:", result.Error)
		return entity.RegistUser{}, result.Error
	}

	dataUser := entity.RegistUser{
		Name:  user.Name,
		Email: user.Email,
	}

	return dataUser, nil
}
