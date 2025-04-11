package data

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Data struct {
	e *xorm.Engine
	Func
}

type Func struct {
	Rule   RuleFunc
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
	d := &Data{}
	d.Func = Func{
		Rule: RuleFunc{
			d: d,
		},
		Server: ServerFunc{
			d: d,
		},
	}
	return d, nil
}
