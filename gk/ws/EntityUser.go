package ws

var (
	UserDao  = &DaoUser{}
	KvDao    = &DaoKeyVal{}
	ScoreLog = &DaoScoreLog{}
	MsgDao   = &DaoMsg{}
	TokenDao = &DaoToken{}
)

type DaoUser struct {
	List func(siteId, s int64) ([]map[string]interface{}, error) `select Id,Username,DauCount,Dau,Score from GkUser where SiteId=? order by DauCount desc limit ?,20`

	Count func(siteId int64) (int64, bool, error) `select count(*) from GkUser where SiteId=? `

	Setting    func(editType, Info interface{}, privacy, UserId int64, SiteId int64) (int64, error) `update GkUser set EditType=?,Info=?,Privacy=? where Id=? and SiteId=?`
	SettingGet func(SiteId int64, UserId int64) (map[string]string, bool, error)                    `select Info,EditType,Privacy from GkUser where SiteId=? and Id=?`

	BaseInfo           func(SiteId int64, UserId interface{}) (map[string]string, bool, error) `select Password,PasswordError,Id,Username,Score,FollowCount,FavCount,MsgCount,TopicCount,EditType,LoginAward from GkUser where SiteId=? and Id=?`
	BaseInfoByUsername func(SiteId int64, Username string) (map[string]string, bool, error)    `select Password,PasswordError,Id,Username,Score,FollowCount,FavCount,MsgCount,TopicCount,EditType,LoginAward from GkUser where SiteId=? and Username=?`

	MemberInfo func(SiteId int64, username string) (map[string]string, bool, error) `select Username,TopicCount,Id,date_format(CreateAt,'%Y-%m-%d %H:%i:%S') CreateAt,Privacy,Info from GkUser where SiteId=? and Username=?`

	SetPasswordError func(PasswordError interface{}, username string, SiteId int64) (int64, error) `update GkUser set PasswordError=?,LoginDate=now() where Username=? and SiteId=?  `

	UpTopic  func(tc, UserId int64, SiteId int64) (int64, error)      `update GkUser set TopicCount=?,ModifyAt=now() where Id=? and SiteId=? `
	UpReply  func(rc, UserId int64, SiteId int64) (int64, error)      `update GkUser set ReplyCount=?,ModifyAt=now() where Id=? and SiteId=? `
	UpFav    func(rc, UserId int64, SiteId int64) (int64, error)      `update GkUser set FavCount=?,ModifyAt=now() where Id=? and SiteId=? `
	UpFollow func(rc, UserId int64, SiteId int64) (int64, error)      `update GkUser set FollowCount=?,ModifyAt=now() where Id=? and SiteId=? `
	UpMsg    func(UserId, UserId2 int64, SiteId int64) (int64, error) `update GkUser set MsgCount=(select count(*) from Msg where UserId=? and status=0) where Id=? and SiteId=? `

	LoginAward func(SiteId int64, userId int64) (int, bool, error)   `select LoginAward from GkUser where SiteId=? and  Id=?`
	Score      func(SiteId int64, userId int64) (int64, bool, error) `select Score from GkUser where SiteId=? and Id=?`

	UpScore      func(sc, userId, siteId int64) (int64, error)   `update GkUser set Score=? where Id=? and SiteId=?`
	LoginAwardDo func(SiteId int64, userId int64) (int64, error) `update GkUser set LoginAward=0 where SiteId=? and Id=? and LoginAward=1`

	Dau      func(SiteId int64, ) ([]map[string]string, error)           `select Dau,Id,Username,Info from GkUser where SiteId=? and  Dau>0 order by Dau desc limit 0,10`
	DauOrder func(SiteId int64, userId interface{}) (int64, bool, error) `select count(*)+1 DauOrder from GkUser where SiteId=? and  Dau>(select Dau from GkUser where Id=?)`
	DauAdd   func(SiteId int64, userId int64) (int64, error)             `update GkUser set Dau=Dau+1 where SiteId=? and Id=?`

	CheckByEmail    func(SiteId int64, email string) (int64, bool, error)    `select count(*) from GkUser where SiteId=? and Email=?`
	CheckByUsername func(SiteId int64, Username string) (int64, bool, error) `select count(*) from GkUser where SiteId=? and Username=?`

	Add func(SiteId int64, Username, Password, Email, Mobile string) (int64, error) `insert into GkUser(SiteId,Username,Password,Email,Mobile,PasswordError,CreateAt,ModifyAt,LastErrorTime) VALUES (?,?,?,?,?,0,now(),now(),now())`
}

type DaoKeyVal struct {
	Add    func(SiteId int64, k string, v interface{}) (int64, error) `insert into KeyVal(SiteId,KeyName,Val)values(?,?,?)`
	Get    func(SiteId int64, key string) (string, bool, error)       `select Val from KeyVal where SiteId=? and KeyName=?`
	Update func(val, key string, SiteId int64) (int64, error)         `update KeyVal set Val=? where KeyName=? and SiteId=?`

	TopicCount func(SiteId, SiteId2 int64, ) (int64, error) `update KeyVal set Val=(SELECT sum(TopicCount) from GkUser where SiteId=? ) where SiteId=? and KeyName='TopicCount'`
	ReplyCount func(SiteId, SiteId2 int64, ) (int64, error) `update KeyVal set Val=(SELECT sum(ReplyCount) from GkUser where SiteId=? ) where SiteId=? and KeyName='ReplyCount'`
	UserCount  func(SiteId, SiteId2 int64, ) (int64, error) `update KeyVal set Val=(SELECT count(*) from GkUser where SiteId=? ) where SiteId=? and KeyName='UserCount'`

	CountInfo func(SiteId int64, ) ([][]string, error) `select KeyName,Val from KeyVal where SiteId=? and KeyName in ('UserCount','ReplyCount','TopicCount')`
}

type DaoScoreLog struct {
	Count func(SiteId int64, UserId int64) (int64, bool, error) `select count(*) from ScoreLog where SiteId=? and UserId=?`

	List func(SiteId int64, UserId, start int64) ([]map[string]string, error) `select Score,ScoreType,ScoreDesc,EntityId,date_format(CreateAt,'%Y-%m-%d %H:%i:%S') CreateAt,Fee from ScoreLog where SiteId =? and UserId=? order by Id desc limit ?,20`

	Add func(SiteId, Score int64, ScoreType, ScoreDesc, EntityId string, Fee, UserId int64) (int64, error) `insert into ScoreLog(SiteId ,Score,ScoreType, ScoreDesc, EntityId, CreateAt,Fee,UserId)VALUES(?,?,?,?,?,now(),?, ?)`
}

type DaoMsg struct {
	Add func(SiteId int64, Title, Body, EntityId, MsgType string, FromUserId, userId, groupId int64) (int64, error) `INSERT INTO Msg(SiteId,Title,Body,EntityId,MsgType,FromUserId,UserId,Status,CreateAt,GroupId)VALUES(?,?,?,?,?,?,?,0,now(),?)`

	FindGroupId func(SiteId int64, UserId, Id int64) (int64, bool, error) `select GroupId from Msg where SiteId=? and userId=? and Id=?`

	Del func(SiteId int64, id, userId int64) (int64, error) `delete from Msg where SiteId=? and Id=? and UserId=?`

	Count       func(SiteId int64, id int64) (int64, bool, error)                   `select count(*) from Msg where SiteId=? and UserId=?`
	CountUnread func(SiteId int64, id int64) (int64, bool, error)                   `select count(*) from Msg where SiteId=? and UserId=? and Status=0`
	CountByType func(SiteId int64, userId int64, mType string) (int64, bool, error) `select count(*) from Msg where SiteId=? and UserId=? and MsgType=?`

	Read      func(SiteId int64, userId, id int64) (int64, error)      `update Msg set Status=1 where SiteId=? and UserId=? and Id=? and Status=0`
	ReadGroup func(SiteId int64, userId, groupId int64) (int64, error) `update Msg set Status=1 where SiteId=? and UserId=? and GroupId=? and Status=0`

	ListUnread func(SiteId int64, userId, start int64) ([]map[string]string, error)                     `select m.Id,m.Status, m.Title,m.Body,m.FromUserId,u.Username,fmt(m.CreateAt) CreateAt from Msg m left join GkUser u on m.FromUserId=u.Id where m.SiteId=? and m.UserId=? and m.Status=0 order by m.Id desc limit ?,20`
	ListByType func(SiteId int64, userId int64, mType string, start int64) ([]map[string]string, error) `select m.Id,m.Status, m.Title,m.Body,m.FromUserId,u.Username,fmt(m.CreateAt) CreateAt from Msg m left join GkUser u on m.FromUserId=u.Id where m.SiteId=? and m.UserId=? and m.MsgType=? order by m.Id desc limit ?,20`

	List func(SiteId int64, userId, start int64) ([]map[string]string, error) `select m.Id,m.Status, m.Title,m.Body,m.FromUserId,u.Username,fmt(m.CreateAt) CreateAt from Msg m left join GkUser u on m.FromUserId=u.Id where  m.SiteId=? and m.UserId=? order by m.Id desc limit ?,20`
}

type DaoToken struct {
	Del  func(ua, userId interface{}, SiteId int64) (int64, error)        `delete from GkToken where Ua=? and UserId=? and SiteId=?`
	Add  func(UserId, Ua, Token interface{}, SiteId int64) (int64, error) `insert into GkToken(UserId,Ua,Token,CreateAt,SiteId)VALUES(?,?,?,now(),?)`
	Find func(SiteId, userId int64, ua string) (string, bool, error)      `select Token from GkToken where SiteId=? and UserId=? and Ua=? `
}
