package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/redshiftkeying/wd6/server/model"
	"github.com/redshiftkeying/wd6/server/utils"
)

// UserName: raw username, telphone or email address
// Password: md5+sal
// PlatformID: admin platform, user platform, ios platform or android platform
// BotCode: Verification Code
type LoginRequest struct {
	UserName   string `json:"username"`
	PassWord   string `json:"password"`
	PlatformID string `json:"platform_id"`
	BotCode    string `json:"code"`
}

// Response format
// response.data{
//	code: int
//	data:{ json_object }
//	msg: string
//}
func AuthenticationHandler() gin.HandlerFunc {
	r := LoginRequest{}
	user := new(model.User)
	null := new(model.User)
	res := new(Response)
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&r, binding.JSON); err == nil {
			data := make(map[string]interface{}, 2)
			// model.Auth
			if (*user).Auth(r.UserName, r.PassWord) {
				res.Code = 0
				res.Msg = "login successsful"
				token, _ := utils.GetToken(user.Name)
				data["user"] = *user
				data["token"] = token
				res.Data = data
				c.JSON(http.StatusOK, res.MakeGinResponse())
			} else {
				res.Code = 401
				data["user"] = *null
				data["token"] = ""
				res.Msg = "Incorrect username or password"
				res.Data = data
				c.JSON(http.StatusOK, res.MakeGinResponse())
			}
		}
	}
}
