package main

import (
	"github.com/Molsbee/mock-server/model"
	"github.com/Molsbee/mock-server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var corsHandler = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
})

func main() {
	adminRouter := gin.Default()
	adminRouter.Use(corsHandler)
	adminRouter.GET("/collections", func(c *gin.Context) {
		collections := service.GetCollections()
		c.JSON(200, collections)
	})
	adminRouter.POST("/collections", func(c *gin.Context) {
		// Create Collection
		var collection model.Collection
		if err := c.BindJSON(&collection); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := service.CreateCollection(collection.Name); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusCreated)
	})
	adminRouter.DELETE("/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) != 0 {
			if err := service.DeleteCollection(collectionName); err != nil {
				c.Status(http.StatusOK)
				return
			}
		}
		// Delete Collection
	})
	adminRouter.GET("/collections/:collectionName", func(c *gin.Context) {
		// Get Collection Details
	})
	adminRouter.POST("/collections/:collectionName", func(c *gin.Context) {
		// Overwrite collection details
	})
	adminRouter.PUT("/collections/:collectionName", func(c *gin.Context) {
		// Add item to collection details
	})

	adminRouter.Run(":8081")

	//router := gin.Default()
	//router.Use(corsHandler)
	//
	//dir, _ := os.Getwd()
	//files, _ := os.ReadDir(dir + "/collections")
	//// TODO: Only supporting files in the directory not working on sub directories
	//for _, file := range files {
	//	var routes []Route
	//	serverBytes, _ := os.ReadFile(dir + "/collections/" + file.Name())
	//	json.Unmarshal(serverBytes, &routes)
	//
	//	for _, route := range routes {
	//		go func(router *gin.Engine, route Route) {
	//			router.Handle(route.Method, route.Path, func(c *gin.Context) {
	//				c.JSON(http.StatusOK, route.Body)
	//			})
	//		}(router, route)
	//	}
	//}
	//
	//router.Run(":8085")
}

type Route struct {
	Path   string
	Method string
	Body   interface{}
}
