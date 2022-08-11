package route

import (
	"crud-test/user"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", user.GetIndexPage)
	r.GET("/member/:Id", user.GetMemberPage)
	r.GET("/addmember", user.GetAddMemberPage)
	r.GET("/deletemember", user.GetDeleteMemberPage)
	r.GET("/updatememberlist", user.GetUpdateMemberListPage)
	r.GET("/updatemember/:Id", user.GetUpdateMemberPage)
	r.POST("/setmember", user.SetMember)
}
