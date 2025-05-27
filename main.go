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
	{ID: "1", Title: "The Thing", Artist: "John Byrne", Issue: "#1", Price: 8.00},
	{ID: "2", Title: "Fantastic Four", Artist: "John Byrne", Issue: "#220", Price: 20.00},
}


// getComics function responds with the list of all comics as formatted JSON. 
// Uses IndentedJSON for readable output.
// Comment from Gin's context.go repo: "IndentedJSON serializes the given struct as pretty JSON (indented + endlines) into the response body."
func getComics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, comics)
}

//postComics adds a comic from JSON received in the request of body.
func postComics(c *gin.Context) {
	var newComic comic

	// Call BindJSON to bind the received JSON to
	// newComic
	if err := c.BindJSON(&newComic); err != nil {
		return``
	}
}






// main sets up the Gin router and defines the GET route for /comics.
// It uses gin.Default() to initialize the router and associates the /comics path with the getComics handler.
func main() {
	router := gin.Default()
	router.GET("/comics", getComics)

	router.Run("localhost:8080")
}

