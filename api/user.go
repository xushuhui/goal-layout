package api

import (
	"fmt"
	"goal-layout/internal/request"
	"goal-layout/internal/services"
	"goal-layout/pkg/core"

	"github.com/gin-gonic/gin"
)

/*

登录

URL:
	/login

POST参数：

	"phone":"2" //手机号
	"password": "1" //密码


返回值：

	"code": 0
	"message": "ok"

*/
func Login(c *gin.Context) {
	var req request.Login

	if e := core.ParseRequest(c, &req); e != nil {

		c.Error(e)
		return
	}
	fmt.Println(">>>", req)
	data, e := services.Login(req)
	if e != nil {
		c.Error(e)
		return
	}
	core.SetData(c, data)
	return

}
