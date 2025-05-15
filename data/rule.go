package data

import (
	"github.com/goccy/go-json"
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
	ListenAddr string `xorm:"varchar(255) notnull"`
	ListenPort int
	TargetType string
	TargetAddr []string
	TargetRule int64
	TargetTag  string
	Config     json.RawMessage `xorm:"json"`
	ServerId   int64           `xorm:"index notnull"`
	CreatedAt  time.Time       `xorm:"created"`
	UpdatedAt  time.Time       `xorm:"updated"`
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
	_, err := r.d.e.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RuleFunc) Delete(nd ...*Rule) error {
	_, err := r.d.e.Delete(nd)
	if err != nil {
		return err
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
