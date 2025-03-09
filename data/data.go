package data

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Data struct {
	e    *xorm.Engine
	Node NodeFunc
}

func New(path string) (*Data, error) {
	e, err := xorm.NewEngine("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("init xorm engine error: %v", err)
	}
	err = e.Sync(new(Node))
	if err != nil {
		return nil, fmt.Errorf("sync tables error: %v", err)
	}
	return &Data{
		e: e,
		Node: NodeFunc{
			Engine: e,
		},
	}, nil
}
