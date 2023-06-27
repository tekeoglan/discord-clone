package main

import (
	"github/tekeoglan/discord-clone/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()

	gin := gin.Default()

	gin.Run()
}
