package handler

import (
	"net/http"
	"score-publisher-svc/pkg/jwtutil"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	//ctx := c.Request.Context()
	user := User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
	}

	token, err := jwtutil.GenerateToken(user.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
