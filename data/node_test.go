package data

import "testing"

func TestNodeFunc_Create(t *testing.T) {
	nd := &Node{
		Name:       "test",
		ListenIP:   "::",
		ListenPort: 8080,
		TargetType: "tcp",
		TargetIP:   "1.1.1.1",
		TargetPort: 80,
	}
	err := d.Node.Create(nd)
	if err != nil {
		t.Error(err)
	}
}
