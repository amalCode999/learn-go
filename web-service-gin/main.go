
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "Ichlasul Amal", Price:56.99},
	{ID: "2", Title: "Blue Train2", Artist: "Ichlasul Amal", Price:56.99},
	{ID: "3", Title: "Blue Train3", Artist: "Ichlasul Amal", Price:56.99},
	{ID: "4", Title: "Blue Train4", Artist: "Ichlasul Amal", Price:56.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8081")
}



