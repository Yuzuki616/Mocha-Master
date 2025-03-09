package data

type Node struct {
	Name       string
	ListenIP   string
	ListenPort int
	TargetType string
	TargetIP   string
	TargetPort int
	Ext        map[string]interface{}
}
