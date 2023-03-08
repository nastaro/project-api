package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nastaro/project-api/database"
	"github.com/nastaro/project-api/routes"
)

func main() {
	r := gin.Default()
	database.ConnectDb()
	routes.ExecuteBookRequest(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
