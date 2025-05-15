package grpc

import (
	"github.com/Yuzuki616/Mocha-Master/data"
	"github.com/goccy/go-json"
	"github.com/orcaman/concurrent-map/v2"
)
import "google.golang.org/grpc"

type Grpc struct {
	d *data.Data
	s *grpc.Server
	c cmap.ConcurrentMap[int64, grpc.ServerStreamingServer[Response]]
	UnimplementedServerServer
}

func (g *Grpc) ListenAndGetRules(request *Request, rsp grpc.ServerStreamingServer[Response]) error {
	ns, err := g.d.Rule.List(request.Id, "")
	if err != nil {
		return err
	}
	rules := make([]*Rule, 0, len(ns))
	for _, n := range ns {
		b, err := json.Marshal(n.Config)
		if err != nil {
			return err
		}
		rules = append(rules, &Rule{
			ServerId:   n.ServerId,
			Name:       n.Name,
			ListenPort: int64(n.ListenPort),
			TargetType: n.TargetType,
			TargetIP:   n.TargetAddr,
			Ext:        b,
		})
	}
	err = rsp.Send(&Response{Rules: rules})
	if err != nil {
		return err
	}
	g.c.Set(request.Id, rsp)
	go func() {
		<-rsp.Context().Done()
		g.c.Remove(request.Id)
	}()
	return nil
}

func (g *Grpc) NotifyRuleChanged(id int64, rules []*Rule) error {
	rsp, ok := g.c.Get(id)
	if !ok {
		return nil
	}
	err := rsp.Send(&Response{Rules: rules})
	if err != nil {
		return err
	}
	return nil
}

func NewGrpc(d *data.Data) *Grpc {
	return &Grpc{
		d: d,
		s: grpc.NewServer(),
	}
}

func (g *Grpc) Start() error {
	RegisterServerServer(g.s, g)
	return nil
}
