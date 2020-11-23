package main

import "github.com/gin-gonic/gin"

func registerRouter(r *gin.Engine) {
	v1Group := r.Group("/api/v1")

	v1Group.GET("/years", ListYearH)
	v1Group.GET("/years/:year/months", ListYearMonthH)
	v1Group.GET("/years/:year/months/:month/:leap/days", ListYearMonthDayH)

	v1Group.GET("/conv/yang-yin/:year/:month/:day", ConvertYang2YinH)
	v1Group.GET("/conv/yin-yang/:year/:month/:leap/:day", ConvertYin2YangH)
}
