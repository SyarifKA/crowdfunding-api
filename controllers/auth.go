package controllers

import (
	"github.com/SyarifKA/crowdfunding-api/dtos"
	"github.com/SyarifKA/crowdfunding-api/lib"
	log "github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/SyarifKA/crowdfunding-api/repository"
	"github.com/gin-gonic/gin"
)

func FindAllUsers(ctx *gin.Context) {
	result, err := repository.FindAllUsers()
	if err != nil {
		log.Debug("error failed to retrieve users")
		return
	}
	lib.HandlerOK(ctx, "Get All Users Success!", result, nil)
	log.Debug("controller get users")
}

func RegistUser(ctx *gin.Context) {
	var form dtos.RegistUser

	if err := ctx.Bind(&form); err != nil {
		lib.HandlerBadReq(ctx, "Invalid input")
		log.Warn("Registration failed due to invalid input.")
		return
	}

	result, err := repository.RegistUser(form)
	if err != nil {
		lib.HandlerBadReq(ctx, "Registration failed")
		log.Error("Registration failed:", err)
		return
	}

	lib.HandlerOK(ctx, "Registration Success!", result, nil)
	log.Debug("Registration success!")
}
