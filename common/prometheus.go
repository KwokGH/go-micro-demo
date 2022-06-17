package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "go-micro.dev/v4/logger"
	"net/http"
	"strconv"
)

func StartPrometheus(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		log.Info("启动成功")
		if err != nil {
			log.Fatal("启动失败")
		}
	}()
}
