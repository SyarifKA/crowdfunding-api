package controllers

import (
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
