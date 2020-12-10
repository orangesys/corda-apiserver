package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//handlerFunc ..
type handlerFunc func(c *gin.Context) interface{}

//response ...
type response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//wrapper ...
func wrapper(handler handlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var res response
		ret := handler(c)
		switch v := ret.(type) {
		case error:
			res = response{
				Msg:  v.Error(),
				Data: nil,
			}
		default:
			res = response{
				Msg:  "",
				Data: v,
			}
		}
		c.JSON(http.StatusOK, res)
	}
}
