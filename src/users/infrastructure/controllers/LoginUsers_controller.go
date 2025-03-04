package controllers

import (
	"chat/src/Users/application"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	loginUser *application.LoginUser
}

func NewLoginUserController(useCase *application.LoginUser) *LoginUserController {
	return &LoginUserController{loginUser: useCase}
}

func (lc *LoginUserController) Run(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

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

	user, err := lc.loginUser.Execute(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso", "user": user})
}
