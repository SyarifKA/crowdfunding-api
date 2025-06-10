package middlewares

import (
	"github.com/SyarifKA/crowdfunding-api/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// func LogMiddlware(c *gin.Context) {
// 	log.RotateLogIfNeeded()

// 	log.Logger.WithFields(logrus.Fields{
// 		"method":  c.Request.Method,
// 		"path":    c.Request.URL.Path,
// 		"ip":      c.ClientIP(),
// 		"headers": c.Request.Header,
// 	}).Info("Incoming request")

// 	c.Next()
// }

func LogWithMessage(message string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"path":    c.FullPath(),
			"ip":      c.ClientIP(),
			"headers": c.Request.Header,
		}).Info(message)

		handler(c)
	}
}
