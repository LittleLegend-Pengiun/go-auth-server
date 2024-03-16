package main

import (
	"jwt-token/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	router(r)
	r.Run() // listen and serve on localhost:3000
}
