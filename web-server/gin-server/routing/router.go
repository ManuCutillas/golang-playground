package routing

import (
	"github.com/gin-gonic/gin"
	md "../middlewares"
	"net/http"
)

func Routing() *gin.Engine {

	// Set router instance
	app := gin.Default()
	app.Use(md.TokenAuthMiddleware())
	app.Use(md.RevisionMiddleware())
	app.Use(md.RequestIdMiddleware())

	// Handle the index route
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "OK"})
	})
	/*
	// Group user related routes together
	user := app.Group("/u")
	{
		user.GET("/login", ensureNotLoggedIn(), showLoginPage)
		user.POST("/login", ensureNotLoggedIn(), performLogin)
		user.GET("/logout", ensureLoggedIn(), logout)
		user.GET("/register", ensureNotLoggedIn(), showRegistrationPage)
		user.POST("/register", ensureNotLoggedIn(), register)
	}

	// Group article related routes together
	articleRoutes := app.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
		articleRoutes.GET("/create", ensureLoggedIn(), showArticleCreationPage)
		articleRoutes.POST("/create", ensureLoggedIn(), createArticle)
	}
	*/
	return app
}

