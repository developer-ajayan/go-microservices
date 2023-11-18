package pkg

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user (validate username and password)

	// For simplicity, assume valid credentials and generate a token
	token, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
