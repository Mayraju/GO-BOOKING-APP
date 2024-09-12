package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayraju/go-booking-app/models"
	"github.com/mayraju/go-booking-app/services"
)

var register services.RegistrationServiceType

type RegisterRouteInterface interface {
	RegistrationRoutes(gin.RouteInfo)
}

type RegisterRouteStruct struct{}

func (r RegisterRouteStruct) RegistrationRoutes(router *gin.Engine) {
	v1 := router.Group("/")
	{
		v1.POST("/Registration", func(ctx *gin.Context) {
			var form models.Registration
			fmt.Println(ctx, "==================")
			error := ctx.BindJSON(&form)
			if error != nil {
				fmt.Printf("Error in Binding values %s", error)
			}
			//age1, err := strconv.ParseInt(age, 10, 64)
			// if err != nil {
			// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Age value must to String"})
			// }
			fmt.Printf("%s,%s,%s,%s", form.FirstName, form.LastName, form.Email, form.Password)

			var data = register.RegistrationService(form)
			ctx.JSON(http.StatusOK, gin.H{"message": data})
		})

		v1.POST("/Login", func(ctx *gin.Context) {
			var email = ctx.PostForm("email")
			var password = ctx.PostForm("password")

			if email == "" || password == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please Provide Correct email and Password"})
			}

			//fmt.Println("Email %s and %s", email, password)
			data, error := register.LoginService(email, password)
			if error != nil {
				ctx.JSON(http.StatusBadRequest, error)
			}
			ctx.JSON(http.StatusOK, data)
		})
	}
}
