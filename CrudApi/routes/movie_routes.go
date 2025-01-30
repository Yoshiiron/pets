package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       int       `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies = []Movie{
	Movie{ID: 1, Isbn: "4312", Title: "Transformers", Director: &Director{Firstname: "Michael", Lastname: "Bae"}},
}

func InitRoutes(r *gin.Engine) {
	route := r.Group("/")
	{
		route.GET("/movies", GetMovies)
		route.GET("/movies/:id", GetMovie)
		route.POST("/movie", CreateMovie)
		route.PUT("/movie/:id", UpdateMovie)
		route.DELETE("/movie/:id", DeleteMovie)
	}
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": movies,
	})
}

func GetMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is invalid",
		})
		return
	}
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is not provided",
		})
		return
	}
	for _, movie := range movies {
		if movie.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"result": movie,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is invalid",
			})
			return
		}
	}
}

func CreateMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindBodyWithJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	movie.ID = len(movies) + 1
	movies = append(movies, movie)
	c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var updatedMovie Movie
	if err := c.ShouldBindBodyWithJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	for index, movie := range movies {
		if movie.ID == id {
			updatedMovie.ID = id
			movies[index] = updatedMovie
			c.JSON(http.StatusOK, gin.H{
				"result": updatedMovie,
			})
			break
		}
	}
}

func DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for index, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
