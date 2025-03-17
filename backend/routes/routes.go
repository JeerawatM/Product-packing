package routes

import (
	"database/sql"
	"go-backend/controllers"
	"go-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	// ================= GET =================
	router.GET("/api/products", func(c *gin.Context) {
		controllers.GetProducts(c, db)
	})
	router.GET("/api/products/:product_id", func(c *gin.Context) {
		controllers.GetProductsByID(c, db)
	})
	router.GET("/api/boxes", func(c *gin.Context) {
		controllers.GetBoxes(c, db)
	})
	router.GET("/api/boxes/:box_id", func(c *gin.Context) {
		controllers.GetBoxesByID(c, db)
	})
	router.GET("/api/history", func(c *gin.Context) {
		controllers.GetHistory(c, db)
	})
	router.GET("/api/history/:id", func(c *gin.Context) {
		controllers.GetHistoryDetail(c, db)
	})
	router.GET("/api/customers", func(c *gin.Context) {
		controllers.GetCustomers(c, db)
	})
	router.GET("/api/customers/:customer_id", func(c *gin.Context) {
		controllers.GetCustomersByID(c, db)
	})
	router.GET("/api/orderdels", func(c *gin.Context) {
		controllers.GetOrderdels(c, db)
	})

	// ================= POST =================
	router.POST("/api/boxes", func(c *gin.Context) {
		controllers.CreateBoxes(c, db)
	})
	router.POST("/api/products", func(c *gin.Context) {
		controllers.CreateProduct(c, db)
	})
	router.POST("/api/orderdels", func(c *gin.Context) {
		controllers.CreateOrderdels(c, db)
	})
	router.POST("/api/generate", func(c *gin.Context) {
		controllers.GenerateProduct(c, db)
	})
	router.POST("/api/customers", func(c *gin.Context) {
		controllers.CreateCustomers(c, db)
	})
	router.POST("/api/user", func(c *gin.Context) {
		controllers.CreateUsers(c, db)
	})

	router.POST("/api/login", controllers.Login(db))

	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "This is a protected route"})
		})
	}

	router.POST("/api/logout", controllers.Logout)

	// ================= DELETE =================
	router.DELETE("/api/products/:product_id", func(c *gin.Context) {
		controllers.DeleteProduct(c, db)
	})
	router.DELETE("/api/orderdels/:order_del_id", func(c *gin.Context) {
		controllers.DeleteOrderDel(c, db)
	})
	router.DELETE("/api/boxes/:box_id", func(c *gin.Context) {
		controllers.DeleteBoxes(c, db)
	})
	router.DELETE("/api/customers/:customer_id", func(c *gin.Context) {
		controllers.DeleteCustomer(c, db)
	})

	// ================= PUT =================
	router.PUT("/api/boxes/:box_id", func(c *gin.Context) {
		controllers.UpdateBoxes(c, db)
	})
	router.PUT("/api/products/:product_id", func(c *gin.Context) {
		controllers.UpdateProduct(c, db)
	})
	router.PUT("/api/history/:history_id", func(c *gin.Context) {
		controllers.UpdateHistory(c, db)
	})
	router.PUT("/api/customers/:customer_id", func(c *gin.Context) {
		controllers.UpdateCustomers(c, db)
	})

	return router
}
