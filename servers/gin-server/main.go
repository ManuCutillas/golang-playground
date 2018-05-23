/**
 *  Gin Server
 *
 *  created_at May 20, 2018
 *  author: Manu Cutillas<manucutillas@outlook.com>
 *  license: MIT
 */
package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "github.com/joho/godotenv"
  "os"
  "net/http"
  r "./routing"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)
	// Set routing
	app := r.Routing()
	setup(app)
}

func setup(engine *gin.Engine) {
	addr := ":" + loadENV()
	if addr == ":" {
		addr = ":3000"
	}
	if err := engine.RunTLS(addr, "./servers/gin-server/ssl/server.pem", "./servers/gin-server/ssl/server.key"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func loadENV() (string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpsPort := os.Getenv("HTTPS_PORT")
	return httpsPort
}





