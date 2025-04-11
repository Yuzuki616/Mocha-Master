package data

import (
	"time"
	"xorm.io/builder"
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

type RuleFunc struct {
	d *Data
}

func (n *RuleFunc) Create(nd ...*Rule) error {
	_, err := n.d.e.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *RuleFunc) Update(nd *Rule) error {
	_, err := n.d.e.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *RuleFunc) Delete(nd ...*Rule) error {
	_, err := n.d.e.Delete(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *RuleFunc) Get(nd *Rule) error {
	_, err := n.d.e.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *RuleFunc) List(serverId int64, targetType string) ([]Rule, error) {
	var nodes []Rule
	if serverId == 0 {
		err := n.d.e.Find(&nodes)
		if err != nil {
			return nil, err
		}
		return nodes, nil
	}
	err := n.d.e.Where(
		builder.Eq{
			"server_id":   serverId,
			"target_type": targetType,
		}).Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (n *RuleFunc) IsExist(nd *Rule) (bool, error) {
	return n.d.e.Exist(nd)
}
