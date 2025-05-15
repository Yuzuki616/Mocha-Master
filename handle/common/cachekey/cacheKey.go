package cachekey

import "strconv"

const (
	serverIPKeyTmpl = "server_ip:%d"
)

func ServerIPKey(serverID int64) string {
	return serverIPKeyTmpl + ":" + strconv.FormatInt(serverID, 10)
}
