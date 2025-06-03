package admin

import (
	"fmt"
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/gin-gonic/gin"
)

type CreateRuleRequest struct {
	ServerId   int64    `json:"server_id" validate:"required"`
	Name       string   `json:"name" validate:"required"`
	ListenAddr string   `json:"listen_addr" validate:"required"`
	TargetAddr []string `json:"target_addr" validate:"required"`
	Config     string   `json:"config"`
}

// CreateRuleHandle creates a new forwarding rule
// @Summary Create rule
// @Description Create a new traffic forwarding rule for a server
// @Tags rules
// @Accept json
// @Produce json
// @Param request body CreateRuleRequest true "Rule creation request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/rule/create [post]
func (h *Handler) CreateRuleHandle(c *gin.Context) {
	var req CreateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	ok, err := h.Data.Server.IsExist(&data.Server{Id: req.ServerId})
	if err != nil {
		c.JSON(500, response.CommonResponse{
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
		ListenAddr: req.ListenAddr,
		TargetAddr: req.TargetAddr,
		Config:     req.Config,
	}
	err = h.Data.Rule.Create(nd)
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

type UpdateRuleRequest struct {
	Id         int64    `json:"id" validate:"required"`
	ServerId   int64    `json:"server_id" validate:"required"`
	Name       string   `json:"name"`
	ListenAddr string   `json:"listen_addr"`
	TargetType string   `json:"target_type"`
	TargetAddr []string `json:"target_ip"`
	Config     string   `json:"config"`
}

// UpdateRuleHandle updates an existing forwarding rule
// @Summary Update rule
// @Description Update an existing traffic forwarding rule
// @Tags rules
// @Accept json
// @Produce json
// @Param request body UpdateRuleRequest true "Rule update request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/rule/update [post]
func (h *Handler) UpdateRuleHandle(c *gin.Context) {
	var req UpdateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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
		ListenAddr: req.ListenAddr,
		TargetType: req.TargetType,
		TargetAddr: req.TargetAddr,
		Config:     req.Config,
	}
	err = h.Data.Rule.Update(nd)
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

type DeleteRuleRequest struct {
	Id int64 `json:"id" validate:"required"`
}

// DeleteRuleHandle deletes a forwarding rule
// @Summary Delete rule
// @Description Delete a traffic forwarding rule by ID
// @Tags rules
// @Accept json
// @Produce json
// @Param request body DeleteRuleRequest true "Rule deletion request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/rule/delete [post]
func (h *Handler) DeleteRuleHandle(c *gin.Context) {
	var req DeleteRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  fmt.Errorf("delete rule error: %w", err).Error(),
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
		c.JSON(400, response.CommonResponse{
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

type ListRuleRequest struct {
	ServerId int64 `json:"server_id" validate:"required"`
}

// ListRuleHandle lists rules for a server
// @Summary List rules
// @Description Get a list of all forwarding rules for a specific server
// @Tags rules
// @Accept json
// @Produce json
// @Param request body ListRuleRequest true "Rule list request"
// @Success 200 {object} response.CommonResponse{data=[]data.Rule}
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/rule/list [post]
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
		Data: nodes,
	})
}
