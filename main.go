package main

import (
	"github.com/adjiemkyn/go-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/antrin", antriancontroller.Index)
	r.GET("/api/antrian/:id", antriancontroller.Show)
	r.POST("/api/antrian", antriancontroller.Create)
	r.PUT("/api/antrian/:id", antriancontroller.Update)
	r.DELETE("/api/antrian", antriancontroller.Delete)

	r.Run()
}
