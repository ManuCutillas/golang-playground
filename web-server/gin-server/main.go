/**
 *  Gin Server
 *
 *  created_at May 20, 2018
 *  author: Manu Cutillas<manucutillas@outlook.com>
 *  license: MIT
 */
package main

import (
	"html/template"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	md "./middlewares"
)

var html = template.Must(template.New("https").Parse(`
	<html>
		<head>
  			<title>Gin Server</title>
		</head>
		<body>
  			<h1>Welcome to Gin Server!</h1>
		</body>
	</html>
`))

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	HTTP_PORT := os.Getenv("HTTP_PORT")
	HTTPS_PORT := os.Getenv("HTTPS_PORT")

	router := gin.Default()
	router.Use(md.RedirectTo)

	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {

		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// HTTP
	go router.Run(":"+HTTP_PORT)

	// Listen and Server in https://127.0.0.1:8080
	router.RunTLS(":"+HTTPS_PORT, "./web-server/gin-server/ssl/server.pem", "./web-server/gin-server/ssl/server.key")
}


