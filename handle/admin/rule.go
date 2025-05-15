package admin

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type CreateRuleRequest struct {
	ServerId   int64           `json:"server_id" validate:"required"`
	Name       string          `json:"name" validate:"required"`
	ListenIP   string          `json:"listen_ip" validate:"required"`
	ListenPort int             `json:"listen_port" validate:"required"`
	TargetIP   []string        `json:"target_ip" validate:"required"`
	TargetPort []int           `json:"target_port" validate:"required"`
	Config     json.RawMessage `json:"ext"`
}

func (h *Handler) CreateRuleHandle(c *gin.Context) {
	var req CreateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	ok, err := h.Data.Server.IsExist(&data.Server{Id: req.ServerId})
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	if !ok {
		c.JSON(400, response.CommonResponse{
			Code: 400,
			Msg:  "server not exist",
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		ServerId:   req.ServerId,
		Name:       req.Name,
		ListenPort: req.ListenPort,
		TargetAddr: req.TargetIP,
		Config:     req.Config,
	}
	err = h.Data.Rule.Create(nd)
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
	})
}

type UpdateRuleRequest struct {
	Id         int64           `json:"id" validate:"required"`
	ServerId   int64           `json:"server_id" validate:"required"`
	Name       string          `json:"name"`
	ListenIP   string          `json:"listen_ip"`
	ListenPort int             `json:"listen_port"`
	TargetType string          `json:"target_type"`
	TargetAddr []string        `json:"target_ip"`
	TargetPort []int           `json:"target_port"`
	Config     json.RawMessage `json:"config"`
}

func (h *Handler) UpdateRuleHandle(c *gin.Context) {
	var req UpdateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		Id:         req.Id,
		ServerId:   req.ServerId,
		Name:       req.Name,
		ListenPort: req.ListenPort,
		TargetType: req.TargetType,
		TargetAddr: req.TargetAddr,
		Config:     req.Config,
	}
	err = h.Data.Rule.Update(nd)
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
	})
}

type DeleteRuleRequest struct {
	Id int64 `json:"id" validate:"required"`
}

func (h *Handler) DeleteRuleHandle(c *gin.Context) {
	var req DeleteRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		Id: req.Id,
	}
	err = h.Data.Rule.Delete(nd)
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
	})
}

func (h *Handler) DeleteRuleByServerIdHandle(c *gin.Context) {
	var req DeleteRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		ServerId: req.Id,
	}
	err = h.Data.Rule.Delete(nd)
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
	})
}

type ListRuleRequest struct {
	ServerId int64 `json:"server_id" validate:"required"`
}

func (h *Handler) ListRuleHandle(c *gin.Context) {
	var req ListRuleRequest
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
