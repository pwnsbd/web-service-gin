package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represent data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums will response with the list of ll album in JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbums album // created a varibale to catch the new album data

	// call bindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbums); err != nil {
		return
	}

	// Add new album to the album slice
	albums = append(albums, newAlbums)
	c.IndentedJSON(http.StatusCreated, albums)
}

/*
	getALbumsByID locates the album whose ID value matches the id parameter sent by client,
	then returns the albums as response
*/

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// iterating over the albums and matching the id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}
