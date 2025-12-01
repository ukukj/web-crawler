// Package consts provides target sites information
package consts

type Site struct {
	Name string
	URL  string
}

var Recruits = []Site{
	{Name: "잡코리아", URL: "https://jobkorea.co.kr"},
	{Name: "사람인", URL: "https://saramin.co.kr"},
}

// Welfare Seoul, Gyeongi youth welfare
var Welfare = []Site{
	{Name: "청년몽땅정보통", URL: "https://youth.seoul.go.kr"},
	{Name: "경기청년포탈", URL: "https://youth.gg.go.kr"},
}
