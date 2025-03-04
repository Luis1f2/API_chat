package main

import (
	"log"

	messageInfra "chat/src/Message/infrastructure"
	userInfra "chat/src/Users/infrastructure"
	"chat/src/core"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := core.GetDBPool()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(cors.Default())

	userInfra.InitUsers(db, router)

	messageInfra.InitMessages(db, router)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
