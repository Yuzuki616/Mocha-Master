package data

import (
	"time"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	TunOutType = "tun_out"
	TunInType  = "tun_in"
)

type Rule struct {
	Id         int64  `xorm:"pk autoincr"`
	Name       string `xorm:"varchar(255) notnull unique"`
	ListenIP   string
	ListenPort int
	TargetType string
	TargetIP   []string
	TargetPort []int
	TargetRule int64
	TargetTag  string
	Ext        map[string]interface{}
	ServerId   int64     `xorm:"index notnull"`
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type NodeFunc struct {
	*xorm.Engine
}

func (n *NodeFunc) Create(nd ...*Rule) error {
	_, err := n.Engine.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Update(nd *Rule) error {
	_, err := n.Engine.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Delete(nd ...*Rule) error {
	_, err := n.Engine.Delete(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Get(nd *Rule) error {
	_, err := n.Engine.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) List(serverId int64, targetType string) ([]Rule, error) {
	var nodes []Rule
	if serverId == 0 {
		err := n.Engine.Find(&nodes)
		if err != nil {
			return nil, err
		}
		return nodes, nil
	}
	err := n.Engine.Where(
		builder.Eq{
			"server_id":   serverId,
			"target_type": targetType,
		}).Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (n *NodeFunc) IsExist(nd *Rule) (bool, error) {
	return n.Engine.Exist(nd)
}
