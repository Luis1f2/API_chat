package main

import (
	"chat/src/core"
	"chat/src/users/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	db, err := core.GetDBPool()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear el router de Gin
	router := gin.Default()

	// Inicializar las dependencias del módulo de usuarios
	infrastructure.InitUsers(db, router)

	// Iniciar el servidor en el puerto 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
