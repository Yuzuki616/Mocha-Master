package middleware

import "github.com/gin-gonic/gin"

func (m *Middleware) ServerAuth(c *gin.Context) {
	if c.Query("token") != m.c.PullToken {
		c.JSON(403, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	c.Next()
}
