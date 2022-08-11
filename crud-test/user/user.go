package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Member struct {
	Id     int
	Name   string
	Nick   string
	Team   string
	Detail string
	Img    string
}

var PArr = []Member{
	{Id: 0, Name: "김찬희", Nick: "조니", Team: "경영팀", Detail: "CEO", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U01BP0D1RL7-bf7f2279c04f-512"},
	{Id: 1, Name: "김병찬", Nick: "딘", Team: "개발팀", Detail: "CTO", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U01GLM8VAQP-9d28a510d6c7-512"},
	{Id: 2, Name: "윤녹두", Nick: "그린", Team: "개발팀", Detail: "Engineer", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U03FQRV2L4E-67f2a9724d33-512"},
	{Id: 3, Name: "이호선", Nick: "밥", Team: "개발팀", Detail: "Engineer", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U03G2J7BECA-9e6367ddbecc-512"},
	{Id: 4, Name: "문성석", Nick: "케빈", Team: "개발팀", Detail: "Engineer", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U03SGS1RQ06-97b9ef6478a7-512"},
	{Id: 5, Name: "황종인", Nick: "피죤", Team: "개발팀", Detail: "Engineer", Img: "https://ca.slack-edge.com/T01C3N3MS1Y-U03SGS1U3GE-d6c2d5a8e97d-512"},
}

func AddMember(p Member) {
	PArr = append(PArr, p)
}

func removeArr(inputArr []Member, idx int) []Member {
	var result []Member
	result = append(inputArr[:idx], inputArr[idx+1:]...)
	return result
}

func addArrIdx(inputArr []Member, p Member, idx int) []Member {
	var result []Member
	var resultTmp []Member
	inputArrTmp := make([]Member, len(inputArr))
	copy(inputArrTmp, inputArr)
	resultTmp = append(inputArrTmp[:idx], p)
	result = append(resultTmp, inputArr[idx:]...)
	return result
}

func GetIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "CRUD sample",
			"payload": PArr,
		},
	)
}

func GetMemberIdx(id int) int {
	for i := 0; i < len(PArr); i++ {
		if PArr[i].Id == id {
			return i
		}
	}
	return -1
}

func GetMember(id int) Member {
	for i := 0; i < len(PArr); i++ {
		if PArr[i].Id == id {
			return PArr[i]
		}
	}
	return Member{}
}

func GetMemberPage(c *gin.Context) {
	if pId, err := strconv.Atoi(c.Param("Id")); err == nil {
		p := GetMember(pId)
		c.HTML(
			http.StatusOK,
			"member.html",
			gin.H{
				"title":   p.Nick + "의 프로필",
				"payload": p,
			},
		)
	}
}

func GetAddMemberPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"addMember.html",
		gin.H{
			"title": "New Member",
		},
	)
}

func GetDeleteMemberPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"deleteMember.html",
		gin.H{
			"title":   "Delete Member",
			"payload": PArr,
		},
	)
}

func GetUpdateMemberListPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"updateMemberList.html",
		gin.H{
			"title":   "Update Member",
			"payload": PArr,
		},
	)
}

func GetUpdateMemberPage(c *gin.Context) {
	if pId, err := strconv.Atoi(c.Param("Id")); err == nil {
		p := GetMember(pId)
		c.HTML(
			http.StatusOK,
			"updateMember.html",
			gin.H{
				"title":   "Update Member",
				"payload": p,
			},
		)
	}
}

func SetMember(c *gin.Context) {
	switch c.PostForm("state") {
	case "a":
		AddMember(Member{len(PArr), c.PostForm("name"), c.PostForm("nick"), c.PostForm("team"), c.PostForm("detail"), c.PostForm("img")})
	case "d":
		idArr := c.PostFormArray("id")
		for i := len(idArr) - 1; i >= 0; i-- {
			pId, _ := strconv.Atoi(idArr[i])
			idx := GetMemberIdx(pId)
			PArr = removeArr(PArr, idx)
		}
	case "u":
		id, _ := strconv.Atoi(c.PostForm("id"))
		idx := GetMemberIdx(id)
		p := Member{id, c.PostForm("name"), c.PostForm("nick"), c.PostForm("team"), c.PostForm("detail"), c.PostForm("img")}
		PArr = removeArr(PArr, idx)
		PArr = addArrIdx(PArr, p, idx)
	}
	c.Redirect(http.StatusFound, "/")
}
