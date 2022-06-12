package serverimp

import (
	"Tiktok/global"
	"strings"
)

type Server_Utils struct {
}

//拼接生成可访问url
func (se *Server_Utils) UrlUnParse(str string) (url string) {
	if str != "" && !strings.HasPrefix(str, "http") && !strings.HasPrefix(str, "https") {
		url = global.GVA_STATICURL + str
	} else {
		return str
	}
	return url
}
