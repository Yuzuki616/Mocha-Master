package server

import (
	c2 "context"
	"github.com/Yuzuki616/Mocha-Master/data"
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

type GetConfigResponse struct {
	*data.Server `json:",inline"`
	Rules        []data.Rule `json:"rules"`
}

// GetConfigHandle retrieves server configuration
// @Summary Get server config
// @Description Get configuration and rules for a specific server
// @Tags server-api
// @Accept json
// @Produce json
// @Param request body GetConfigRequest true "Server config request"
// @Success 200 {object} response.CommonResponse{data=[]data.Rule}
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /server/getConfig [get]
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
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  "internal server error",
			Data: nil,
		})
		c.Error(err)
		return
	}
	server := &data.Server{Id: req.ServerId}
	err = h.Data.Server.Get(server)
	if err != nil {
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  "internal server error",
			Data: nil,
		})
		c.Error(err)
		return
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: GetConfigResponse{
			Server: server,
			Rules:  nodes,
		},
	})
}

type ReportStatusRequest struct {
	ServerId int64    `json:"server_id" validate:"required"`
	IP       []net.IP `json:"ip" validate:"required"`
}

// ReportStatusHandle reports server status
// @Summary Report server status
// @Description Report server status including IP addresses
// @Tags server-api
// @Accept json
// @Produce json
// @Param request body ReportStatusRequest true "Server status report"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /server/reportStatus [post]
func (h *Handler) ReportStatusHandle(c *gin.Context) {
	var req ReportStatusRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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
			Msg:  "internal server error",
			Data: nil,
		})
		c.Error(err)
		return
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}
