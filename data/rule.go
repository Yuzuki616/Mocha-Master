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
	Id         int64     `xorm:"pk autoincr" json:"id"`
	Name       string    `xorm:"varchar(255) notnull unique" json:"name"`
	ListenAddr string    `xorm:"varchar(255) notnull" json:"listen_addr"`
	TargetType string    `xorm:"varchar(255)" json:"target_type"`
	TargetAddr []string  `xorm:"json" json:"target_addr"`
	TargetRule int64     `xorm:"index" json:"target_rule"`
	TargetTag  string    `xorm:"varchar(255)" json:"target_tag"`
	Config     string    `json:"config"`
	ServerId   int64     `xorm:"index notnull" json:"server_id"`
	CreatedAt  time.Time `xorm:"created" json:"created_at"`
	UpdatedAt  time.Time `xorm:"updated" json:"updated_at"`
}

type RuleFunc struct {
	d *Data
}

func (r *RuleFunc) Create(nd ...*Rule) error {
	_, err := r.d.e.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) Update(nd *Rule) error {
	_, err := r.d.e.ID(nd.Id).Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) Delete(nd ...*Rule) error {
	for _, n := range nd {
		if n.Id == 0 {
			_, err := r.d.e.Delete(n)
			if err != nil {
				return err
			}
		} else {
			_, err := r.d.e.ID(n.Id).Delete(n)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RuleFunc) Get(nd *Rule) error {
	_, err := r.d.e.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) List(serverId int64, targetType string) ([]Rule, error) {
	var nodes []Rule
	if serverId == 0 {
		err := r.d.e.Find(&nodes)
		if err != nil {
			return nil, err
		}
		return nodes, nil
	}
	err := r.d.e.Where(
		builder.Eq{
			"server_id":   serverId,
			"target_type": targetType,
		}).Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (r *RuleFunc) IsExist(nd *Rule) (bool, error) {
	return r.d.e.Exist(nd)
}
