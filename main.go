package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mayraju/go-booking-app/loaders"
	"github.com/mayraju/go-booking-app/routes"
)

var (
	router = gin.Default()
)

func main() {
	fmt.Printf("In main function===============")
	loaders.ConnectDB()
	routes.FCNRoutes(router)
	fmt.Printf("after db connection=========")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
	// })

	router.Run(":9000")
}
