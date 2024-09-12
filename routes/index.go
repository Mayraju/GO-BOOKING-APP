package routes

import (
	"github.com/gin-gonic/gin"
)

var registerRoutes RegisterRouteStruct

func FCNRoutes(router *gin.Engine) {
	registerRoutes.RegistrationRoutes(router)

}
