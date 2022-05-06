package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"mc-server/config"
	"mc-server/mojang"
	"net/http"
)

func Authenticate(c *gin.Context) {
	parseData, _ := c.Get("body")
	parse := parseData.(gjson.Result)
	username := parse.Get("username").String()
	clientToken := parse.Get("clientToken").String()
	uid, ok := config.UserInfo[username]
	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}
	result := gin.H{
		"accessToken": uid,
		"clientToken": clientToken,
		"selectedProfile": gin.H{
			"id":   uid,
			"name": username,
		},
		"availableProfiles": []gin.H{
			{
				"id":   uid,
				"name": username,
			},
		},
	}
	if parse.Get("requestUser").Bool() {
		result["user"] = mojang.UserInfo{
			ID: uid,
			Properties: []mojang.UserInfoProperty{
				{
					Name:  "preferredLanguage",
					Value: "zh_CN",
				},
			},
		}
	}
	c.JSON(200, result)
}
func ReAuthenticate(c *gin.Context) {
	parseData, _ := c.Get("body")
	parse := parseData.(gjson.Result)
	uid := parse.Get("accessToken").String()
	clientToken := parse.Get("clientToken").String()
	username := ""
	for k, v := range config.UserInfo {
		if v == uid {
			username = k
			break
		}
	}
	if username == "" {
		c.Status(http.StatusForbidden)
		return
	}
	result := gin.H{
		"accessToken": uid,
		"clientToken": clientToken,
		"selectedProfile": gin.H{
			"id":   uid,
			"name": username,
		},
	}
	if parse.Get("requestUser").Bool() {
		result["user"] = mojang.UserInfo{
			ID: uid,
			Properties: []mojang.UserInfoProperty{
				{
					Name:  "preferredLanguage",
					Value: "zh_CN",
				},
			},
		}
	}
	c.JSON(200, result)
}
func MProfile(c *gin.Context) {
	parseData, _ := c.Get("body")
	parse := parseData.(gjson.Result)
	client := mojang.New()
	var usernames []string
	for _, v := range parse.Array() {
		usernames = append(usernames, v.String())
	}
	multipleUUIDs, err := client.FetchMultipleUUIDs(usernames)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	result := make([]mojang.ProfileBase, len(multipleUUIDs))
	for k, v := range multipleUUIDs {
		result = append(result, mojang.ProfileBase{
			UUID: v,
			Name: k,
		})
	}
	c.JSON(200, result)
}

func HasJoined(c *gin.Context) {
	name, _ := c.GetQuery("username")
	uid, ok := config.UserInfo[name]
	if !ok {
		c.Status(http.StatusNoContent)
		return
	}
	client := mojang.New()
	profile, err := client.FetchProfile(uid, false)
	if err != nil {
		c.JSON(200, mojang.ProfileBase{
			UUID: uid,
			Name: name,
		})
		return
	}
	c.JSON(200, profile)
}
func Profile(c *gin.Context) {
	uid := c.Param("uuid")
	name := ""
	for k, v := range config.UserInfo {
		if v == uid {
			name = k
			break
		}
	}
	if name == "" {
		c.Status(http.StatusNoContent)
		return
	}
	client := mojang.New()
	unsigned := true
	unsignedStr, b := c.GetQuery("unsigned")
	if b && unsignedStr == "false" {
		unsigned = false
	}
	profile, err := client.FetchProfile(uid, unsigned)
	if err != nil {
		c.JSON(200, mojang.ProfileBase{
			UUID: uid,
			Name: name,
		})
		return
	}
	c.JSON(200, profile)
}
