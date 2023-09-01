package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var corsHandler = cors.New(cors.Config{
	AllowAllOrigins:  true,
	AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "HEAD"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{},
	AllowCredentials: true,
})

func main() {
	router := gin.Default()
	router.Use(corsHandler)

	dir, _ := os.Getwd()
	files, _ := os.ReadDir(dir + "/collections")
	// TODO: Only supporting files in the directory not working on sub directories
	for _, file := range files {
		var routes []Route
		serverBytes, _ := os.ReadFile(dir + "/collections/" + file.Name())
		json.Unmarshal(serverBytes, &routes)

		for _, route := range routes {
			go func(router *gin.Engine, route Route) {
				router.Handle(route.Method, route.Path, func(c *gin.Context) {
					c.JSON(http.StatusOK, route.Body)
				})
			}(router, route)
		}
	}

	router.Run(":8085")
}

type Route struct {
	Path   string
	Method string
	Body   interface{}
}
