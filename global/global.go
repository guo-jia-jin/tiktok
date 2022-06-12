package global

import (
	"Tiktok/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB      //gorm连接数据库
	GVA_CONFIG    config.Server //全局服务配置
	GVA_VP        *viper.Viper  //viper读取yaml配置文件
	GVA_LOCALIP   string        //静态资源访问不到的下下之策，本机ip
	GVA_STATICURL string        //静态资源路径前缀
)
