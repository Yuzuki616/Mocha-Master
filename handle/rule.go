package handle

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/gin-gonic/gin"
)

type RuleHandler struct {
	Handle
}
type CreateRuleRequest struct {
	ServerId   int64                  `json:"server_id" validate:"required"`
	Name       string                 `json:"name" validate:"required"`
	ListenIP   string                 `json:"listen_ip" validate:"required"`
	ListenPort int                    `json:"listen_port" validate:"required"`
	TargetType string                 `json:"target_type" validate:"required"`
	TargetIP   string                 `json:"target_ip" validate:"required"`
	TargetPort int                    `json:"target_port" validate:"required"`
	Ext        map[string]interface{} `json:"ext"`
}

func (h *RuleHandler) Create(c *gin.Context) {
	var req CreateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	ok, err := h.d.Server.IsExist(&data.Server{Id: req.ServerId})
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	if !ok {
		c.JSON(400, CommonResponse{
			Code: 400,
			Msg:  "server not exist",
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		ServerId:   req.ServerId,
		Name:       req.Name,
		ListenIP:   req.ListenIP,
		ListenPort: req.ListenPort,
		TargetType: req.TargetType,
		TargetIP:   req.TargetIP,
		TargetPort: req.TargetPort,
		Ext:        req.Ext,
	}
	err = h.d.Rule.Create(nd)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}

type UpdateRuleRequest struct {
	Id         int64                  `json:"id" validate:"required"`
	Name       string                 `json:"name"`
	ListenIP   string                 `json:"listen_ip"`
	ListenPort int                    `json:"listen_port"`
	TargetType string                 `json:"target_type"`
	TargetIP   string                 `json:"target_ip"`
	TargetPort int                    `json:"target_port"`
	Ext        map[string]interface{} `json:"ext"`
}

func (h *RuleHandler) Update(c *gin.Context) {
	var req UpdateRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		Id:         req.Id,
		Name:       req.Name,
		ListenIP:   req.ListenIP,
		ListenPort: req.ListenPort,
		TargetType: req.TargetType,
		TargetIP:   req.TargetIP,
		TargetPort: req.TargetPort,
		Ext:        req.Ext,
	}
	err = h.d.Rule.Update(nd)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}

type DeleteRuleRequest struct {
	Id int64 `json:"id" validate:"required"`
}

func (h *RuleHandler) Delete(c *gin.Context) {
	var req DeleteRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		Id: req.Id,
	}
	err = h.d.Rule.Delete(nd)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}

func (h *RuleHandler) DeleteByServerId(c *gin.Context) {
	var req DeleteRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Rule{
		ServerId: req.Id,
	}
	err = h.d.Rule.Delete(nd)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}

type ListRuleRequest struct {
	ServerId int64 `json:"server_id" validate:"required"`
}

func (h *RuleHandler) List(c *gin.Context) {
	nodes, err := h.d.Rule.List()
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: nodes,
	})
}
