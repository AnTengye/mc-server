package main

import (
	"github.com/gin-gonic/gin"
	"mc-server/config"
	"mc-server/handler"
	"mc-server/middlerware"
	"mc-server/mojang"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(middlerware.ParamsPrint())
	r.POST("/authserver/authenticate", handler.Authenticate)
	r.POST("/authserver/refresh", handler.ReAuthenticate)

	r.POST("/sessionserver/session/minecraft/join", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	r.POST("/authserver/validate", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	r.POST("/authserver/invalidate", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	r.POST("/authserver/signout", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	r.GET("/sessionserver/session/minecraft/hasJoined", handler.HasJoined)
	r.GET("/sessionserver/session/minecraft/profile/:uuid", handler.Profile)
	r.POST("/api/profiles/minecraft", handler.MProfile)
	meta := mojang.MetaInfo{
		Meta: map[string]interface{}{
			"serverName":              "antengye",
			"implementationName":      "Yggdrasil API for antengye", //服务端实现的名称
			"implementationVersion":   "1.0.0",                      //服务端实现的版本
			"feature.non_email_login": true,                         //指示验证服务器是否支持使用邮箱之外的凭证登录
			"description":             "play with my friend",
		},
		SkinDomains:        []string{},
		SignaturePublicKey: config.PublicKey,
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, meta)
	})
	r.Run(":8899")
}
