package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// comic represents data about a comic book. !! This represents the kind of information that can be stored/tracked by this API. Similar to database schema.
type comic struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Issue  string  `json:"issue"`
	Writer string  `json:"writer"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// comics is a hardcoded slice of comic data used for testing.
// Each entry matches the fields defined in the comic struct.
var comics = []comic{
	{ID: "1", Title: "The Thing", Writer: "John Byrne", Artist: "John Byrne", Issue: "#1", Price: 8.00},
	{ID: "2", Title: "Fantastic Four", Writer: "John Byrne", Artist: "John Byrne", Issue: "#220", Price: 20.00},
}


// !! <----------------------------------------- ROUTES -------------------------------------------->

// <-------------------------------------------- ROUTE: GET ------------------------------------->
// getComics function responds with the list of all comics as formatted JSON. 
// Uses IndentedJSON for readable output.
// Comment from Gin's context.go repo: "IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body."
func getComics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, comics)
}

// <-------------------------------------------- ROUTE: POST ------------------------------------->
//postComics adds a comic from JSON received in the request of body.
	 {
	var newComic comic

	// Call BindJSON to bind the received JSON to
	// newComic
	if err := c.BindJSON(&newComic); err != nil {
		return
	}

	//Add the new comic to the slice.
	comics = append(comics, newComic)
	c.IndentedJSON(http.StatusCreated, newComic)
}

// <------------------------------------------ ROUTE: GET BY ID ------------------------------------->
// getComicByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getComicByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of comics, looking for
	// a comic whose ID value matches the parameter.
	for _, a := range comics {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "comic not found"})
}


// <------------------------------------------ START SERVER --------------------------------------->
// main sets up the Gin router and defines the GET route for /comics.
// It uses gin.Default() to initialize the router and associates the /comics path with the getComics handler.
func main() {
	router := gin.Default()
	router.GET("/comics", getComics)
	router.GET("/comics/:id", getComicByID)
	router.POST("/comics", postComics)

	router.Run("localhost:8080")
}
