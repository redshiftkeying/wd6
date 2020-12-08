package controller

import (
	"github.com/gin-gonic/gin"
)

// Response modle
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// MakeGinResponse return gin json response object
func (r *Response) MakeGinResponse() gin.H {
	return gin.H{"code": r.Code, "data": r.Data, "msg": r.Msg}
}
