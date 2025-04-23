package handler

import (
	"net/http"

	"github.com/Molsbee/mock-server/model"
	"github.com/Molsbee/mock-server/service"
	"github.com/gin-gonic/gin"
)

type collectionHandler struct {
	repo service.CollectionRepo
}

func NewCollectionHandler(repo service.CollectionRepo) *collectionHandler {
	return &collectionHandler{
		repo: repo,
	}
}

func (h *collectionHandler) GetCollections() (string, string, func(c *gin.Context)) {
	return "GET", "/collections", func(c *gin.Context) {
		collections := h.repo.GetCollectionNames()
		c.JSON(200, collections)
	}
}

func (h *collectionHandler) CreateCollection() (string, string, func(c *gin.Context)) {
	return "POST", "/collections", func(c *gin.Context) {
		var collection model.Collection
		if err := c.BindJSON(&collection); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := h.repo.CreateCollection(collection); err != nil {
			c.JSON(http.StatusInternalServerError, model.APIErrorResponse{Error: err.Error()})
			return
		}
		c.Status(http.StatusCreated)
	}
}

func (h *collectionHandler) DeleteCollection() (string, string, func(c *gin.Context)) {
	return "DELETE", "/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, model.APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		if err := h.repo.DeleteCollection(collectionName); err != nil {
			c.JSON(http.StatusInternalServerError, model.APIErrorResponse{Error: err.Error()})
			return
		}
		c.Status(http.StatusAccepted)
		return
	}
}

func (h *collectionHandler) GetCollection() (string, string, func(c *gin.Context)) {
	return "GET", "/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, model.APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		collection, err := h.repo.GetCollection(collectionName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.APIErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, collection)
		return
	}
}

func (h *collectionHandler) UpdateCollection() (string, string, func(c *gin.Context)) {
	return "PUT", "/collections/:collectionName", func(c *gin.Context) {
		collectionName := c.Param("collectionName")
		if len(collectionName) == 0 {
			c.JSON(http.StatusBadRequest, model.APIErrorResponse{Error: "Please provide a collection name"})
			return
		}

		var collection model.Collection
		if err := c.BindJSON(&collection); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		update, err := h.repo.UpdateCollection(collectionName, collection)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.APIErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, update)
		return
	}
}
