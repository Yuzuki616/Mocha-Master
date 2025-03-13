package data

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Data struct {
	e      *xorm.Engine
	Rule   NodeFunc
	Server ServerFunc
}

func New(path string) (*Data, error) {
	e, err := xorm.NewEngine("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("init xorm engine error: %v", err)
	}
	err = e.Sync(new(Rule), new(Server))
	if err != nil {
		return nil, fmt.Errorf("sync tables error: %v", err)
	}
	return &Data{
		e: e,
		Rule: NodeFunc{
			Engine: e,
		},
		Server: ServerFunc{
			Engine: e,
		},
	}, nil
}
