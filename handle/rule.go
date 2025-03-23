package handle

import (
	"fmt"
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
	TargetIP   []string               `json:"target_ip" validate:"required"`
	TargetPort []int                  `json:"target_port" validate:"required"`
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
	TargetIP   []string               `json:"target_ip"`
	TargetPort []int                  `json:"target_port"`
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
	var req ListRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nodes, err := h.d.Rule.List(req.ServerId, "")
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

type CreateTunRuleRequest struct {
	ServerId       int64                  `json:"server_id" validate:"required"`
	Name           string                 `json:"name" validate:"required"`
	ListenIP       string                 `json:"listen_ip" validate:"required"`
	ListenPort     int                    `json:"listen_port" validate:"required"`
	TargetListenIp string                 `json:"target_listen_ip" validate:"required"`
	TargetId       int64                  `json:"target_id" validate:"required"`
	TargetPort     []int                  `json:"target_port" validate:"required"`
	OutIp          []string               `json:"out_ip" validate:"required"`
	OutPort        []int                  `json:"out_port" validate:"required"`
	Ext            map[string]interface{} `json:"ext"`
}

func (h *RuleHandler) CreateTun(c *gin.Context) {
	var req CreateTunRuleRequest
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
	// Get target server
	ts := &data.Server{Id: req.TargetId}
	err = h.d.Server.Get(ts)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	targetTag := fmt.Sprintf(
		"%s-%d-%d",
		ts.Name,
		ts.Id,
		req.TargetPort[0])

	// Create out rule
	nd2 := &data.Rule{
		ServerId:   req.TargetId,
		Name:       req.Name,
		ListenIP:   req.TargetListenIp,
		ListenPort: req.TargetPort[0],
		TargetIP:   req.OutIp,
		Ext:        req.Ext,
	}
	// Create in rule
	nd := &data.Rule{
		ServerId:   req.ServerId,
		Name:       req.Name,
		ListenIP:   req.ListenIP,
		ListenPort: req.ListenPort,
		TargetIP:   ts.Ip,
		TargetPort: req.TargetPort,
		TargetTag:  targetTag,
		Ext:        req.Ext,
	}

	err = h.d.Rule.CreateTun(nd, nd2)

	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
	})
}

type UpdateTunRuleRequest struct {
	Id             int64                  `json:"id" validate:"required"`
	Name           string                 `json:"name" validate:"required"`
	ListenIP       string                 `json:"listen_ip" validate:"required"`
	ListenPort     int                    `json:"listen_port" validate:"required"`
	TargetRule     int64                  `json:"target_id" validate:"required"`
	TargetListenIp string                 `json:"target_listen_ip" validate:"required"`
	TargetPort     []int                  `json:"target_port" validate:"required"`
	OutIp          []string               `json:"out_ip" validate:"required"`
	OutPort        []int                  `json:"out_port" validate:"required"`
	Ext            map[string]interface{} `json:"ext"`
}

func (h *RuleHandler) UpdateTun(c *gin.Context) {
	var req UpdateTunRuleRequest
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
		TargetPort: req.TargetPort,
		Ext:        req.Ext,
	}

	nd2 := &data.Rule{
		Id:         req.TargetRule,
		TargetType: data.TunOutType,
		Name:       req.Name,
		ListenIP:   req.TargetListenIp,
		ListenPort: req.TargetPort[0],
		TargetIP:   req.OutIp,
		TargetPort: req.OutPort,
		Ext:        req.Ext,
	}
	err = h.d.Rule.UpdateTun(nd, nd2)
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

type DeleteTunRuleRequest struct {
	Id         int64 `json:"id" validate:"required"`
	TargetRule int64 `json:"target_rule" validate:"required"`
}

func (h *RuleHandler) DeleteTun(c *gin.Context) {
	var req DeleteTunRuleRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	err = h.d.Rule.Delete(
		&data.Rule{
			Id: req.Id,
		},
		&data.Rule{
			Id: req.TargetRule,
		})
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
