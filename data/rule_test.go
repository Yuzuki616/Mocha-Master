package data

import "testing"

func TestNodeFunc_Create(t *testing.T) {
	nd := &Rule{
		Name:       "test",
		ListenIP:   "::",
		ListenPort: 8080,
		TargetType: "tcp",
		TargetIP:   "1.1.1.1",
		ServerId:   1,
		TargetPort: 80,
	}
	ok, err := d.Server.IsExist(&Server{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		d.Server.Create(&Server{Id: 1, Name: "test"})
	}
	err = d.Rule.Create(nd)
	if err != nil {
		t.Error(err)
	}
}
