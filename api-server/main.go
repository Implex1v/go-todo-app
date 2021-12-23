package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gokurs/config"
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
