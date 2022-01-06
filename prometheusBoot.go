package common

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	http.Handle("/metric", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		log.Info(err)
		if err != nil {
			log.Fatal("启动失败")
		}
	}()
}
