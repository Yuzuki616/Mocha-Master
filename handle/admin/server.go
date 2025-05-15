package admin

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle/common/cachekey"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net"
)

type ServerRequest struct {
	Name   string          `json:"name"`
	Config json.RawMessage `json:"ext"`
}

func (h *Handler) CreateServerHandle(c *gin.Context) {
	var req ServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Name:   req.Name,
		Config: req.Config,
	}
	err = h.Data.Server.Create(nd)
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

func (h *Handler) UpdateServerHandle(c *gin.Context) {
	var req ServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Name:   req.Name,
		Config: req.Config,
	}
	err = h.Data.Server.Update(nd)
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

type DeleteServerRequest struct {
	Id int64 `json:"id"`
}

func (h *Handler) DeleteServerHandle(c *gin.Context) {
	var req DeleteServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Id: req.Id,
	}
	err = h.Data.Server.Delete(nd)
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

type GetServerRequest struct {
	Id          int64 `json:"id"`
	ContainRule bool  `json:"contain_rule"`
}

type GetServerResponseData struct {
	*data.Server `json:",inline"`
	IP           []net.IP    `json:"ip"`
	Rules        []data.Rule `json:"rules"`
}

func (h *Handler) GetServerHandle(c *gin.Context) {
	var req GetServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Id: req.Id,
	}
	err = h.Data.Server.Get(nd)
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	var rules []data.Rule
	if req.ContainRule {
		rules, err = h.Data.Rule.List(req.Id, "")
		if err != nil {
			c.JSON(200, response.CommonResponse{
				Code: 500,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		return
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: &GetServerResponseData{
			Server: nd,
			Rules:  rules,
		},
	})
}

type ListServerResponse struct {
	data.Server `json:",inline"`
	IP          []net.IP `json:"ip"`
}

func (h *Handler) ListServerHandle(c *gin.Context) {
	nodes, err := h.Data.Server.List()
	if err != nil {
		c.JSON(200, response.CommonResponse{
			Code: 500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	ss := make([]ListServerResponse, 0, len(nodes))
	for _, s := range nodes {
		ip, err := h.Cache.Get(c, cachekey.ServerIPKey(s.Id))
		if err != nil {
			c.JSON(500, response.CommonResponse{
				Code: 500,
				Msg:  err.Error(),
				Data: nil,
			})
			return
		}
		ss = append(ss, ListServerResponse{
			Server: s,
			IP:     ip.([]net.IP),
		})
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: nodes,
	})
}
