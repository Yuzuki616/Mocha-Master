package data

import (
	"time"
	"xorm.io/xorm"
)

type Node struct {
	Name       string `xorm:"varchar(255) notnull pk unique"`
	ListenIP   string
	ListenPort int
	TargetType string
	TargetIP   string
	TargetPort int
	Ext        map[string]interface{}
	CreatedAt  time.Time `xorm:"created"`
	UpdatedAt  time.Time `xorm:"updated"`
}

type NodeFunc struct {
	*xorm.Engine
}

func (n *NodeFunc) Create(nd *Node) error {
	_, err := n.Engine.Insert(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Update(nd *Node) error {
	_, err := n.Engine.Update(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Delete(nd *Node) error {
	_, err := n.Engine.Delete(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) Get(nd *Node) error {
	_, err := n.Engine.Get(nd)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeFunc) List() ([]Node, error) {
	var nodes []Node
	err := n.Engine.Find(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (n *NodeFunc) IsExist(nd *Node) (bool, error) {
	return n.Engine.Exist(nd)
}
