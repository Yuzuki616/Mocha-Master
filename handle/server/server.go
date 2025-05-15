package server

import (
	c2 "context"
	"github.com/Yuzuki616/Mocha-Master/handle/common/cachekey"
	c "github.com/Yuzuki616/Mocha-Master/handle/common/context"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/gin-gonic/gin"
	"net"
	"time"
)

type Handler struct {
	*c.Context
}

type GetConfigRequest struct {
	ServerId int64 `json:"server_id" validate:"required"`
}

func (h *Handler) GetConfigHandle(c *gin.Context) {
	var req GetConfigRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nodes, err := h.Data.Rule.List(req.ServerId, "")
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: nodes,
	})
}

type ReportStatusRequest struct {
	ServerId int64    `json:"server_id" validate:"required"`
	IP       []net.IP `json:"ip" validate:"required"`
}

func (h *Handler) ReportStatusHandle(c *gin.Context) {
	var req ReportStatusRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	err = h.Cache.Set(
		c2.Background(),
		cachekey.ServerIPKey(req.ServerId),
		req.IP,
		store.WithExpiration(time.Minute*5))
	if err != nil {
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}
