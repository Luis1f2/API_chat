package main

import (
	"log"

	messageInfra "chat/src/Message/infrastructure"
	userInfra "chat/src/Users/infrastructure"
	"chat/src/core"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	db, err := core.GetDBPool()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close() // Defer justo después de abrir el recurso

	// Crear el router de Gin
	router := gin.Default()

	userInfra.InitUsers(db, router)

	messageInfra.InitMessages(db, router)

	// Iniciar el servidor en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
