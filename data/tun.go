package data

import "fmt"

type CreateTunParams struct {
	ServerId       int64
	Name           string
	ListenIP       string
	ListenPort     int
	TargetListenIp string
	TargetId       int64
	OutIp          []string
	OutPort        []int
	Ext            map[string]interface{}
}

func (r *RuleFunc) CreateTun(p *CreateTunParams) error {
	// Get target server
	ts := &Server{Id: p.TargetId}
	err := r.d.Server.Get(ts)
	if err != nil {
		return err
	}
	targetTag := fmt.Sprintf(
		"%s-%d",
		ts.Name,
		ts.Id)

	// Create out rule
	nd2 := &Rule{
		ServerId:   p.TargetId,
		Name:       p.Name,
		ListenIP:   p.TargetListenIp,
		ListenPort: p.ListenPort,
		TargetIP:   p.OutIp,
		TargetPort: p.OutPort,
		Ext:        p.Ext,
	}
	// Create in rule
	nd := &Rule{
		ServerId:   p.ServerId,
		Name:       p.Name,
		ListenIP:   p.ListenIP,
		ListenPort: p.ListenPort,
		TargetIP:   ts.Ip,
		TargetPort: []int{p.ListenPort},
		TargetTag:  targetTag,
		Ext:        p.Ext,
	}
	// Create out rule
	err = r.Create(nd2)
	if err != nil {
		return err
	}
	// Create in rule
	nd.TargetRule = nd2.Id
	nd.TargetType = TunInType
	nd2.TargetType = TunOutType
	err = r.Create(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) UpdateTun(nd *Rule, nd2 *Rule) error {
	// Update out rule
	err := r.Update(nd2)
	if err != nil {
		return err
	}
	// Update in rule
	nd.TargetRule = nd2.Id
	nd.TargetType = TunInType
	nd2.TargetType = TunOutType
	err = r.Update(nd)
	if err != nil {
		return err
	}
	return nil
}
