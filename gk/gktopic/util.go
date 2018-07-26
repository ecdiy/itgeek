package gktopic

func topicLink(topicId, userId, title string) string {
	return ` <a onclick="vgo('/p/topic/detail,` + topicId + `,` + userId + `',this)">` + title + `</a> `
}
func memberLink(un string) string {
	return ` <a onclick="vgo('/p/member/` + un + `')">` + un + `</a> `
}
