package admin

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/Yuzuki616/Mocha-Master/handle/common/cachekey"
	"github.com/Yuzuki616/Mocha-Master/handle/common/response"
	"github.com/gin-gonic/gin"
	"net"
)

type ServerRequest struct {
	Name      string `json:"name" validate:"required"`
	PortRange [2]int `json:"port_range" validate:"required"`
	Config    string `json:"config" validate:"required json"`
}

// CreateServerHandle creates a new server
// @Summary Create server
// @Description Create a new server with the provided configuration
// @Tags servers
// @Accept json
// @Produce json
// @Param request body ServerRequest true "Server creation request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/server/create [post]
func (h *Handler) CreateServerHandle(c *gin.Context) {
	var req ServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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

type UpdateServerRequest struct {
	Id     int64  `json:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Config string `json:"config" validate:"required json"`
}

// UpdateServerHandle updates an existing server
// @Summary Update server
// @Description Update an existing server with the provided configuration
// @Tags servers
// @Accept json
// @Produce json
// @Param request body UpdateServerRequest true "Server update request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/server/update [post]
func (h *Handler) UpdateServerHandle(c *gin.Context) {
	var req UpdateServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	nd := &data.Server{
		Id:     req.Id,
		Name:   req.Name,
		Config: req.Config,
	}
	println(req.Id)
	err = h.Data.Server.Update(nd)
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

type DeleteServerRequest struct {
	Id int64 `json:"id"`
}

// DeleteServerHandle deletes a server
// @Summary Delete server
// @Description Delete a server by ID
// @Tags servers
// @Accept json
// @Produce json
// @Param request body DeleteServerRequest true "Server deletion request"
// @Success 200 {object} response.CommonResponse
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/server/delete [post]
func (h *Handler) DeleteServerHandle(c *gin.Context) {
	var req DeleteServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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

type GetServerRequest struct {
	Id          int64 `json:"id"`
	ContainRule bool  `json:"contain_rule"`
}

type GetServerResponseData struct {
	*data.Server `json:",inline"`
	IP           []net.IP    `json:"ip"`
	Rules        []data.Rule `json:"rules"`
}

// GetServerHandle gets server details
// @Summary Get server
// @Description Get server details by ID, optionally including rules
// @Tags servers
// @Accept json
// @Produce json
// @Param request body GetServerRequest true "Server get request"
// @Success 200 {object} response.CommonResponse{data=GetServerResponseData}
// @Failure 400 {object} response.CommonResponse
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/server/get [post]
func (h *Handler) GetServerHandle(c *gin.Context) {
	var req GetServerRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, response.CommonResponse{
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
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  "internal server error",
			Data: nil,
		})
		c.Error(err)
		return
	}
	var rules []data.Rule
	if req.ContainRule {
		rules, err = h.Data.Rule.List(req.Id, "")
		if err != nil {
			c.JSON(500, response.CommonResponse{
				Code: 500,
				Msg:  "internal server error",
				Data: nil,
			})
			c.Error(err)
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

// ListServerHandle lists all servers
// @Summary List servers
// @Description Get a list of all servers with their IP addresses
// @Tags servers
// @Produce json
// @Success 200 {object} response.CommonResponse{data=[]ListServerResponse}
// @Failure 500 {object} response.CommonResponse
// @Security ApiKeyAuth
// @Router /admin/server/list [get]
func (h *Handler) ListServerHandle(c *gin.Context) {
	nodes, err := h.Data.Server.List()
	if err != nil {
		c.JSON(500, response.CommonResponse{
			Code: 500,
			Msg:  "internal server error",
			Data: nil,
		})
		c.Error(err)
		return
	}
	ss := make([]ListServerResponse, 0, len(nodes))
	for _, s := range nodes {
		var ips []net.IP
		ip, _ := h.Cache.Get(c, cachekey.ServerIPKey(s.Id))
		if ip != nil {
			ips = ip.([]net.IP)
		} else {
			ips = make([]net.IP, 0)
		}
		ss = append(ss, ListServerResponse{
			Server: s,
			IP:     ips,
		})
	}
	c.JSON(200, response.CommonResponse{
		Code: 200,
		Msg:  "success",
		Data: nodes,
	})
}
