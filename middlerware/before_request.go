package middlerware

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
)

func ParamsPrint() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			rawData, _ := c.GetRawData()
			parse := gjson.ParseBytes(rawData)
			log.Printf("%+v", parse.String())
			c.Set("body", parse)
		} else {
			log.Printf("%+v", c.Params)
		}
		c.Next()
	}
}
