package service

import serverimp "Tiktok/service/serverImp"

type ServerGroup struct {
	Servers serverimp.ServerGroup
}

var ServerGroupApp = new(ServerGroup)
