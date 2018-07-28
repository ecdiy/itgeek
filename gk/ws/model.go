package ws

type UpReq struct {
	UserId int64
	Type   string
	Val    int64
	// 积分变化值
	Fee       int64
	EntityId  string
	ScoreType string
	ScoreDesc string
	//
	SiteId int64
}
type MsgReq struct {
	UserId          int64
	FromUserId      int64
	Title           string
	Body            string
	SiteId, GroupId int64
	EntityId        string
	MsgType         string
	FromUsername    string
}
