package handler

import (
	"Tiktok/handler/api"
)

//API集合
//向外提供所有API
type ApiGroup struct {
	ApiGroup api.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
