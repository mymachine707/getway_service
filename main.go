package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"mymachine707/clients"
	"mymachine707/config"
	docs "mymachine707/docs" // docs is generated by Swag CLI, you have to import it.
	"mymachine707/handlars"

	_ "github.com/lib/pq"
)

// @license.name	Apache 2.0
func main() {

	cfg := config.Load()
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	fmt.Println("----->>")
	fmt.Printf("%+v\n", cfg)
	fmt.Println("---->>")

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery
	}

	r.GET("/ping", MyCORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	defer grpcClients.Close()

	h := handlars.NewHandler(cfg, grpcClients)
	v1 := r.Group("v1")
	{
		v1.Use(MyCORSMiddleware())
		v1.POST("/login", h.Login)
		v1.POST("/client", h.CreatClient)
		v1.GET("/client/:id", h.AuthMiddleware("*"), h.GetClientByID)
		v1.GET("/client", h.AuthMiddleware("sudo"), h.GetClientList)
		v1.PUT("/client", h.AuthMiddleware("*"), h.ClientUpdate)
		v1.DELETE("/client/:id", h.AuthMiddleware("sudo"), h.DeleteClient)

		v1.POST("/product", h.AuthMiddleware("sudo"), h.CreatProduct)
		v1.GET("/product/:id", h.AuthMiddleware("sudo"), h.GetProductByID)
		v1.GET("/product", h.AuthMiddleware("*"), h.GetProductList)
		v1.PUT("/product", h.AuthMiddleware("sudo"), h.ProductUpdate)
		v1.DELETE("/product/:id", h.AuthMiddleware("sudo"), h.DeleteProduct)
		v1.GET("/my-product/:id", h.AuthMiddleware("sudo"), h.SearchProductByMyUsername)

		v1.POST("/category", h.AuthMiddleware("sudo"), h.CreatCategory)
		v1.GET("/category/:id", h.AuthMiddleware("sudo"), h.GetCategoryByID)
		v1.GET("/category", h.AuthMiddleware("*"), h.GetCategoryList)
		v1.PUT("/category", h.AuthMiddleware("sudo"), h.CategoryUpdate)
		v1.DELETE("/category/:id", h.AuthMiddleware("sudo"), h.DeleteCategory)

		v1.POST("/order", h.AuthMiddleware("*"), h.CreatOrder)
		v1.GET("/order/:id", h.AuthMiddleware("*"), h.GetOrderByID)
		v1.GET("/order", h.AuthMiddleware("*"), h.GetOrderList)
		v1.PUT("/order", h.AuthMiddleware("*"), h.OrderUpdate)
		v1.DELETE("/order/:id", h.AuthMiddleware("*"), h.DeleteOrder)

		v1.POST("/orderItem", h.AuthMiddleware("*"), h.CreatOrderItem)
		v1.GET("/orderItem/:id", h.AuthMiddleware("*"), h.GetOrderItemByID)
		v1.GET("/orderItem", h.AuthMiddleware("*"), h.GetOrderItemList)
		v1.DELETE("/orderItem/:id", h.AuthMiddleware("*"), h.DeleteOrderItem)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func MyCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MyCORSMiddleware...")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}
