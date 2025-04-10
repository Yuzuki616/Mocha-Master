package data

import "testing"

func TestRuleFunc_CreateTun(t *testing.T) {
	t.Log(d.Rule.CreateTun(&Rule{
		Name:       "test3",
		ListenIP:   "::",
		ListenPort: 8080,
		ServerId:   1,
	}, &Rule{
		Name:       "test1",
		ListenIP:   "::",
		ListenPort: 8080,
		ServerId:   1,
	}))
}
