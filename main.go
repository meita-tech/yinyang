package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/meita-tech/yinyang/calendar"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/zhangjie2012/cbl-go"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

var (
	Version     string = ""
	Branch      string = ""
	BuildTime   string = ""
	rawDataPath        = "/data/data/rawdata"
	host               = "localhost"
	port               = 9123
)

func main() {
	logrus.Infof("build info: branch=%s, version=%s, buildtime=%s", Branch, Version, BuildTime)

	flag.StringVar(&rawDataPath, "rawdata", rawDataPath, "the yin-yang relate rawdata file path")
	flag.StringVar(&host, "host", host, "the server host")
	flag.IntVar(&port, "port", port, "the server port")
	flag.Parse()

	calendar.ParseRawData(rawDataPath)
	logrus.Infof("parse raw data success")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cbl.PromGinMiddleware())
	registerRouter(router)

	router.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{})))
	if err := router.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
		logrus.Fatalf("server run failure|%s", err)
	}
}
