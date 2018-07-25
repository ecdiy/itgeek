package ws

var (
	RefererDao = &DaoReferer{}
)

type DaoReferer struct {
	Add func(SiteId, TopicId int64, Referer string) (int64, error) `INSERT INTO TopicReferer(SiteId,TopicId,Referer,Click,CreateAt,ModifyAt,ClickAll)VALUES(?,?,?,1,now(),now(),1)`

	Check func(SiteId, TopicId int64, Referer string) (int64, bool, error) `select count(*) from TopicReferer where SiteId=? and TopicId=? and Referer=?`

	Up func(SiteId, TopicId int64, Referer string) (int64, error) `update TopicReferer set Click=Click+1 ,ClickAll=ClickAll+1 where SiteId=? and TopicId=? and Referer=?`
}
