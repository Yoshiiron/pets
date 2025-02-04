package routes

import "github.com/gin-gonic/gin"

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.POST("/book/", controllers.CreateBook)
	r.GET("/book/", controllers.GetBook)
	r.GET("/book/{bookId}", controllers.GetBookById)
	r.PUT("/book/{bookId}", controllers.UpdateBook)
	r.DELETE("/book/{bookId}", controllers.DeleteBook)
}
