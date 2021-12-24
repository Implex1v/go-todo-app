package main

import (
	"api-server/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	c := config.GetConfig()
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DisableConsoleColor()

	router := gin.Default()
	_ = router.Run(fmt.Sprintf(":%s", c.Port))
}
