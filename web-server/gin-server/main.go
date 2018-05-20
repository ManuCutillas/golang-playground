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
	"net/http"
	"os/signal"
	"context"
	"time"
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

	// Server in https://127.0.0.1:8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := router.RunTLS(":"+HTTPS_PORT, "./web-server/gin-server/ssl/server.pem", "./web-server/gin-server/ssl/server.key");
		err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()


	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}


