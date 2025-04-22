package main

import (
	"fmt"
	"net/http"

	"github.com/Molsbee/mock-server/model"
	"github.com/Molsbee/mock-server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var corsHandler = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
})

func setupMockServer() {
	collections := service.GetCollections()
	if len(collections) != 0 {
		router := gin.Default()
		router.Use(corsHandler)
		// Setup Mock Server
		for _, collection := range collections {
			for _, route := range collection.Routes {
				go func(router *gin.Engine, collectionName string, route model.Route) {
					router.Handle(route.Method, fmt.Sprintf("%s/%s", collectionName, route.Path), func(c *gin.Context) {
						c.JSON(route.StatusCode, route.Body)
					})
				}(router, collection.Name, route)
			}
		}
		go router.Run(":8085")
	}
}

func main() {
	setupMockServer()

	adminRouter := gin.Default()
	adminRouter.Use(corsHandler)
	adminRouter.GET("/collections", func(c *gin.Context) {
		collections := service.GetCollectionNames()
		c.JSON(200, collections)
	})
	adminRouter.POST("/collections", func(c *gin.Context) {
		var collection model.Collection
		if err := c.BindJSON(&collection); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := service.CreateCollection(collection); err != nil {
			c.JSON(http.StatusInternalServerError, APIErrorResponse{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
	})
	adminRouter.DELETE("/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		if err := service.DeleteCollection(collectionName); err != nil {
			c.JSON(http.StatusInternalServerError, APIErrorResponse{Error: err.Error()})
			return
		}
		c.Status(http.StatusAccepted)
		return
	})
	adminRouter.GET("/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		collection, err := service.GetCollection(collectionName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, collection)
		return
	})
	adminRouter.PUT("/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		var collection model.Collection
		if err := c.BindJSON(&collection); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		update, err := service.UpdateCollection(collectionName, collection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, APIErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, update)
		return
	})
	adminRouter.Run(":8081")
}

type APIErrorResponse struct {
	Error string `json:"error"`
}
