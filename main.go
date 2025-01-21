package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/revathiasina/ecommerce-go/controllers"
	"github.com/revathiasina/ecommerce-go/databases"
	"github.com/revathiasina/ecommerce-go/middleware"
	"github.com/revathiasina/ecommerce-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(databases.ProductData(databases.Client, "Products"), databases.UserData(databases.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckot", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
