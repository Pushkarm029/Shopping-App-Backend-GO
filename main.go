package main

import (
	"amazon-backend/config"
	sellerauth "amazon-backend/handlers/authentication/seller_auth"
	userauth "amazon-backend/handlers/authentication/user_auth"
	sellerhandler "amazon-backend/handlers/seller_handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()
	// for i := 0; i < 100; i++ {
	// 	user := models.User{}
	// 	err := faker.FakeData(&user)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	config.DB.Save(&user)
	// }
	router := gin.Default()

	router.Use(cors.Default())

	initRoutes(router)

	router.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	r.GET("/api/products", sellerhandler.Get_Product)
	r.POST("/api/user/login", userauth.Login)
	r.POST("/api/user/register", userauth.Register)
	r.POST("/api/seller/login", sellerauth.Login)
	r.POST("/api/seller/register", sellerauth.Register)
	r.POST("/api/seller/addproduct", sellerhandler.AddProduct)
}
