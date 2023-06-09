package main

import (
	"encoding/json"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/silentiumNoxe/buffalo/config"
	"log"
	"net/http"
	"os"
)

func main() {
	var engine = gin.Default()
	gin.Logger()
	var dir = static.LocalFile("./public", false)

	engine.Use(static.Serve("/", dir))
	engine.NoRoute(notFound(dir))

	var serve = func(endpoint, file string) {
		engine.GET(
			endpoint, func(c *gin.Context) {
				c.FileFromFS(file, dir)
			},
		)
	}

	for endpoint, file := range config.Default.Route {
		serve(endpoint, file)
	}

	log.Println("Start server on " + config.Default.GetAddr())
	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func notFound(dir http.FileSystem) gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Default.NotFound == "" {
			return
		}

		c.FileFromFS(config.Default.NotFound, dir)
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
}

func init() {
	payload, err := os.ReadFile("gosfs.json")
	if err != nil {
		log.Println("Config file not found. Used default configuration")
		return
	}

	var c config.Config
	if err := json.Unmarshal(payload, &c); err != nil {
		log.Println("Failed read config - " + err.Error())
		os.Exit(1)
		return
	}

	config.Default.Merge(&c)
}
