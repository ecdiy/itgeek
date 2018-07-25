package ws

var (
	ReplyDao = &DaoReply{}
	ThankDao = &DaoReplyThank{}
)

type DaoReply struct {
	Get func(SiteId, replyId int64) (map[string]string, bool, error) `select r.TopicId,r.UserId,r.Username,t.Title,t.UserId Author from TopicReply r left join Topic t on r.TopicId=t.Id where r.SiteId=? and r.Id=?`

	CountByUserId func(SiteId int64, id interface{}) (int64, bool, error) `select count(*) from TopicReply where SiteId=? and UserId=?`
	Count         func(SiteId int64, TopicId int64) (int64, bool, error)  `select count(*) from TopicReply where SiteId=? and TopicId=?`

	List func(SiteId int64, topicId interface{}) ([]map[string]string, error) `select UserId,Id,Reply,Username,Reply,fmt(CreateAt) CreateAt from TopicReply where SiteId=? and TopicId=? order by id limit 0,200`

	Insert func(SiteId int64, UserId int64, TopicId, Username, Reply interface{}) (int64, error) `INSERT INTO TopicReply(SiteId,UserId,TopicId,Username,Reply,CreateAt)VALUES(?,?,?,?,?, now())`
}

type DaoReplyThank struct {
	List func(SiteId, TopicId, UserId int64) ([]int64, error) `select ReplyId from ReplyThank where SiteId=? and TopicId=? and UserId=?`

	Exist func(SiteId, ReplyId int64) (int, bool, error) `select count(*) from ReplyThank where SiteId=? and ReplyId=?`

	Add func(SiteId, UserId, ReplyId int64, TopicId interface{}) (int64, error) `INSERT INTO ReplyThank(SiteId,UserId,ReplyId,TopicId,CreateAt) VALUES (?, ?,?,?,now())`
}
