package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) UserAuth(c *gin.Context) {
	if c.Query("token") != m.c.AccessToken {
		c.JSON(403, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
