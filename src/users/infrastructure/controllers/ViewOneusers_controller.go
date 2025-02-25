package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"chat/src/users/application"

	"github.com/gin-gonic/gin"
)

type ViewOneUserController struct {
	userViewer *application.ViewUser
}

func NewViewOneUserController(useCase *application.ViewUser) *ViewOneUserController {
	return &ViewOneUserController{userViewer: useCase}
}

func (vc *ViewOneUserController) Run(c *gin.Context) {
	// Convertir ID de string a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Ejecutar caso de uso para obtener usuario
	user, err := vc.userViewer.Execute(id)
	if err != nil {
		// Si el usuario no existe, devolver código 404
		if strings.Contains(err.Error(), "no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		// Otro error interno
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}

	// Asegurar que no se devuelva la contraseña
	response := gin.H{
		"id":       user.ID,
		"username": user.Username,
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}
