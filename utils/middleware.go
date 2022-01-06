package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Di taroh di default / use / new, berlaku untuk semuanya
func GetToken(c *gin.Context) {
	// token := c.Request.Header["Token"] // Harus token
	// stringToken := strings.Join(token, "")

	token := c.GetHeader("token")
	// if stringToken != "cobadulu" {
	fmt.Println(token)
	if token != "cobadulu" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"StatusCode": http.StatusUnauthorized,
			"Message":    "Unauthorize",
		})
		return
	}
}
