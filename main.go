package main

import (
	"polling/polling/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterRouter(r)

	// Puerto en el que se va a escuchar el servidor
	port := ":8081"

	// Levantamos el servidor
	r.Run(port)

}
