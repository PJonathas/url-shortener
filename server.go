package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pjonathas/url-shortener/adapters/generator"
	"github.com/pjonathas/url-shortener/adapters/storage"
	"github.com/pjonathas/url-shortener/shortener"
)

var svcStorage shortener.Storage
var svcGenerator shortener.Generator
var shortenerSvc shortener.Service

func init() {
	svcStorage = storage.Mongo{}
	svcGenerator = generator.Default{}
	shortenerSvc = shortener.Service{Storage: svcStorage, Generator: svcGenerator}
}

func main() {
	r := gin.Default()

	r.POST("/shorten", shorten)

	r.POST("/unshorten", unshorten)

	r.Run()
}

func shorten(c *gin.Context) {
	var json struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := shortenerSvc.Shorten(json.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func unshorten(c *gin.Context) {
	var json struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := shortenerSvc.Unshorten(json.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
