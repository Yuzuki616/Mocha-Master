package handle

import "github.com/gin-gonic/gin"

type UserHandler struct {
	Handle *Handle
}

func (h *UserHandler) TokenCheck(c *gin.Context) {
	c.JSON(200, &CommonResponse{
		Code: 200,
		Msg:  "success",
	})
	return
}
