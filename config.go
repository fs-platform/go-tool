package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀
		consul.WithPrefix(prefix),
		//是否移除前缀,true:表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	con, err := config.NewConfig()
	if err != nil {
		return con, err
	}
	//加载配置
	err = con.Load(consulSource)
	return con, err
}
