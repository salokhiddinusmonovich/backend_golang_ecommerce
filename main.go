package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/salokhiddinusmonovich/backend_golang_ecommerce/controller"
	"github.com/salokhiddinusmonovich/backend_golang_ecommerce/database"
	"github.com/salokhiddinusmonovich/backend_golang_ecommerce/midlleware"
	"github.com/salokhiddinusmonovich/backend_golang_ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controller.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.User(gin.Logger())

	routes.UserRoutes(router)
	router.Use(midlleware.Authetication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
