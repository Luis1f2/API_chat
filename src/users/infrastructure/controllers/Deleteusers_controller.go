package controllers

import (
	"chat/src/Users/application"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	userRemover *application.DeleteUser
}

func NewDeleteUserController(useCase *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{userRemover: useCase}
}

func (uc *DeleteUserController) Run(c *gin.Context) {
	// Convertir ID de string a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Ejecutar caso de uso para eliminar usuario
	err = uc.userRemover.Execute(id)
	if err != nil {
		// Si el usuario no existe, devolver código 404
		if strings.Contains(err.Error(), "no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		// Otro error interno
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
