package controllers

import (
	"chat/src/Users/application"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SaveUserController struct {
	userSaver *application.SaveUser
}

func NewSaveUserController(useCase *application.SaveUser) *SaveUserController {
	return &SaveUserController{userSaver: useCase}
}

func (uc *SaveUserController) Run(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Validar si el JSON recibido es válido
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	body.Username = strings.TrimSpace(body.Username)
	body.Password = strings.TrimSpace(body.Password)

	if body.Username == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de usuario y la contraseña son obligatorios"})
		return
	}

	if len(body.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La contraseña debe tener al menos 6 caracteres"})
		return
	}

	err := uc.userSaver.Execute(body.Username, body.Password)
	if err != nil {

		if strings.Contains(err.Error(), "usuario ya existe") {
			c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya existe"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario guardado exitosamente"})
}
