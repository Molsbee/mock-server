package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Molsbee/mock-server/handler"
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

func setupMockServer(repo service.CollectionRepo) *http.Server {
	router := gin.Default()
	router.Use(corsHandler)

	collections := repo.GetCollections()
	if len(collections) != 0 {
		// Setup handlers for Mock Server
		for _, collection := range collections {
			for _, route := range collection.Routes {
				go func(router *gin.Engine, collectionName string, route model.Route) {
					router.Handle(route.Method, fmt.Sprintf("%s/%s", collectionName, route.Path), func(c *gin.Context) {
						for key, value := range route.Headers {
							c.Header(key, value)
						}
						c.JSON(route.StatusCode, route.Body)
					})
				}(router, collection.Name, route)
			}
		}
	}

	// Using server to support shutting it down programatically
	srv := &http.Server{
		Addr:    ":8085",
		Handler: router,
	}
	go srv.ListenAndServe()

	return srv
}

func main() {
	repo := service.NewFileCollectionRepo()
	mockServer := setupMockServer(repo)

	adminRouter := gin.Default()
	adminRouter.Use(corsHandler)

	collectionHandler := handler.NewCollectionHandler(repo)
	adminRouter.Handle(collectionHandler.GetCollections())
	adminRouter.Handle(collectionHandler.CreateCollection())
	adminRouter.Handle(collectionHandler.DeleteCollection())
	adminRouter.Handle(collectionHandler.GetCollection())
	adminRouter.Handle(collectionHandler.UpdateCollection())
	adminRouter.POST("/server/restart", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := mockServer.Shutdown(ctx)
		if err == nil {
			mockServer = setupMockServer(repo)
		}
	})
	adminRouter.Run(":8081")
}
