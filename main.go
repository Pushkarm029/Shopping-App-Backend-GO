// Amazon Features :

//You'll need to install various Go packages to set up your GraphQL API.
// Popular packages include github.com/graphql-go/graphql, github.com/graphql-go/handler,
// github.com/jinzhu/gorm, and github.com/lib/pq.
// You can install them using go get

// User Side Features ->
// User can add items in cart
// Remove items from cart
// Users can view purchased items
// Users can sign-in. -> for seller or customer
// Users can cancel orders.
// Users can return or exchange items.
// Users can search for products.
// Users can place orders.

// Seller Side Features ->
// Inventory management -> Check out of stock
// add or remove product

// Procedure ->
// PostgreSQL -> init with some fake data
// Data Stucture -> products
// —------------------------------------------------------------------------------------------

// Tech Stack
// Go-Gin
// PostgreSQL
// NextTS or ReactTS

package main

import (
	"amazon-backend/config"
	sellerhandler "amazon-backend/handlers/seller_handler"
	"amazon-backend/models"
	"fmt"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	var users []models.User
	for i := 0; i < 10; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, user)
	}
	r := gin.Default()
	r.Use(cors.Default())
	initRoutes(r)
	fmt.Print(users)
	router.Run(":8080")
}

func initRoutes(r *gin.Engine) {
	r.GET("/api/users", sellerhandler.Get_Product)
}


