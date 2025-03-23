package handle

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/gin-gonic/gin"
)

type ServerHandler struct {
	Handle
}

type ServerRequest struct {
	Name string                 `json:"name"`
	Ext  map[string]interface{} `json:"ext"`
}

func (h *ServerHandler) Create(c *gin.Context) {
	var req ServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Name: req.Name,
		Ext:  req.Ext,
	}
	err = h.d.Server.Create(nd)
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

func (h *ServerHandler) Update(c *gin.Context) {
	var req ServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Name: req.Name,
		Ext:  req.Ext,
	}
	err = h.d.Server.Update(nd)
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

type DeleteServerRequest struct {
	Id int64 `json:"id"`
}

func (h *ServerHandler) Delete(c *gin.Context) {
	var req DeleteServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Id: req.Id,
	}
	err = h.d.Server.Delete(nd)
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

type GetServerRequest struct {
	Id          int64 `json:"id"`
	ContainRule bool  `json:"contain_rule"`
}

type GetServerResponseData struct {
	Server *data.Server `json:"server"`
	Rules  []data.Rule  `json:"rules"`
}

func (h *ServerHandler) Get(c *gin.Context) {
	var req GetServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Id: req.Id,
	}
	err = h.d.Server.Get(nd)
	if err != nil {
		c.JSON(200, CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	var rules []data.Rule
	if req.ContainRule {
		rules, err = h.d.Rule.List(req.Id, "")
		if err != nil {
			c.JSON(200, CommonResponse{
				Code: 500,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		return
	}
	c.JSON(200, CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: &GetServerResponseData{
			Server: nd,
			Rules:  rules,
		},
	})
}

func (h *ServerHandler) List(c *gin.Context) {
	nodes, err := h.d.Server.List()
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
