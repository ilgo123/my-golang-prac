package main

import(
	"github.com/ilgo123/go-ecommerce/controllers"
	"github.com/ilgo123/go-ecommerce/database"
	"github.com/ilgo123/go-ecommerce/middleware"
	"github.com/ilgo123/go-ecommerce/routes"
	"github.com/gin-gionic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = 8000
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addcart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}