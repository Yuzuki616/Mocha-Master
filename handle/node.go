package handle

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/gin-gonic/gin"
)

type NodeHandler struct {
	Handle
}

type NodeRequest struct {
	Name       string                 `json:"name"`
	ListenIP   string                 `json:"listen_ip"`
	ListenPort int                    `json:"listen_port"`
	TargetType string                 `json:"target_type"`
	TargetIP   string                 `json:"target_ip"`
	TargetPort int                    `json:"target_port"`
	Ext        map[string]interface{} `json:"ext"`
}

func (h *NodeHandler) Create(c *gin.Context) {
	var req NodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Node{
		Name:       req.Name,
		ListenIP:   req.ListenIP,
		ListenPort: req.ListenPort,
		TargetType: req.TargetType,
		TargetIP:   req.TargetIP,
		TargetPort: req.TargetPort,
		Ext:        req.Ext,
	}
	err = h.d.Node.Create(nd)
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

func (h *NodeHandler) Update(c *gin.Context) {
	var req NodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Node{
		Name:       req.Name,
		ListenIP:   req.ListenIP,
		ListenPort: req.ListenPort,
		TargetType: req.TargetType,
		TargetIP:   req.TargetIP,
		TargetPort: req.TargetPort,
		Ext:        req.Ext,
	}
	err = h.d.Node.Update(nd)
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

type DeleteNodeRequest struct {
	Name string `json:"name"`
}

func (h *NodeHandler) Delete(c *gin.Context) {
	var req DeleteNodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Node{
		Name: req.Name,
	}
	err = h.d.Node.Delete(nd)
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

type GetNodeRequest struct {
	Name string `json:"name"`
}

func (h *NodeHandler) Get(c *gin.Context) {
	var req GetNodeRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Node{
		Name: req.Name,
	}
	err = h.d.Node.Get(nd)
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
		Data: nd,
	})
}

func (h *NodeHandler) List(c *gin.Context) {
	nodes, err := h.d.Node.List()
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
