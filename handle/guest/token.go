package guest

import (
	"github.com/Yuzuki616/Mocha-Master/handle/common/context"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*context.Context
}

type TokenRequest struct {
	Token string `json:"token"`
}

// TokenCheck validates access token
// @Summary Check token
// @Description Validate access token for authentication
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body TokenRequest true "Token validation request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 403 {object} response.CommonResponse
// @Router /tokenCheck [post]
func (h *Handler) TokenCheck(c *gin.Context) {
	var req TokenRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, &response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}
	if h.Conf.AccessToken == req.Token {
		c.JSON(200, &response.CommonResponse{
			Code: 200,
			Msg:  "success",
		})
		return
	}
	c.JSON(403, &response.CommonResponse{
		Code: 403,
		Msg:  "forbidden",
	})
	return
}
