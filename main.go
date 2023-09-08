package main

import (
	"gin-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.InitRouters(r)
	r.Run(":8080")
}
