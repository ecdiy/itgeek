package ws

var (
	TopicDao         = &DaoTopic{}
	TopicCategoryDao = &DaoTopicCategory{}
	FavDao           = &DaoFav{}
	FollowDao        = &DaoFollow{}
	AppendDao        = &DaoTopicAppend{}
)

type DaoTopicCategory struct {
	List func(SiteId int64) ([]map[string]string, error) `select Id,Name,ParentId,ItemCount from Category where SiteId=?`

	Del func(SiteId, Id int64) (int64, error) `delete from Category where SiteId=? and Id=?`

	Add func(SiteId int64, Name, ParentId interface{}) (int64, error) `INSERT INTO  Category (SiteId,Name,ParentId,CreateAt,ModifyAt,ItemCount) VALUES (?,?, ?, now(), now(),0)`

	UpItemCount func(SiteId int64, id interface{}) (int64, error) `update Category set ItemCount=ItemCount+1 where SiteId=? and Id=?`

	UpName func(name string, Id, SiteId int64) (int64, error) `update Category set Name=? where Id=? and SiteId=? `
}

type DaoTopic struct {
	FindBase func(SiteId, Id int64) (map[string]string, bool, error) `select Title,UserId,Id from Topic where SiteId=? and Id=?`

	UpCategory func(toId, catId, siteId int64) (int64, error) `update Topic set CategoryId=? where CategoryId=? and SiteId=?`

	Find          func(id interface{}) (map[string]string, bool, error)   `select Id,Title,UserId from Topic where Id=?`
	Count         func(SiteId int64, ) (int64, bool, error)               `select count(*) from Topic where SiteId=?`
	CountByUserId func(SiteId int64, id interface{}) (int64, bool, error) `select count(*) from Topic where SiteId=? and UserId=?`

	UpdateReply func(ReplyUsername string, ReplyCount, ReplyUserId, Id int64) (int64, error) `update Topic set ReplyUsername=?,ReplyCount=?,ReplyUserId=?,ReplyTime=now() where Id=?`

	Add func(SiteId int64, Title, Body, UserId, CategoryId, Username, SourceType, Source interface{}) (int64, error) `INSERT INTO Topic(SiteId,Title,Body,UserId,CategoryId,CreateAt,Username,SourceType,Source)VALUES(?,?,?,?,?,now(),?,?,?)`

	ModifyShowTimes func(SiteId int64, Id, UserId interface{}) (int64, error) `update Topic set Today=Today+1,ShowTimes=ShowTimes+1 where SiteId=? and Id=? and UserId=?`

	FindById func(SiteId int64, Id, UserId interface{}) (map[string]string, bool, error) ` select n.*,c.Name ParentCatName from (
				select t.UserId, t.Title,t.Username,t.Body,t.ReplyCount,t.ShowTimes,t.ReplyUsername,t.ReplyUserId,fmt(t.CreateAt) CreateAt,t.CategoryId,c.Name,c.ParentId																					
             from Topic t left join Category c on c.Id=t.CategoryId where t.SiteId=? and t.Id=? and t.UserId=?) n left join Category c on n.ParentId=c.Id `

	List func(SiteId int64, start int64) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt)CreateAt,c.Name CategoryName,
		t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime from Topic t left join Category c on t.CategoryId=c.Id 
							where t.SiteId=?  order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`

	Top1000 func(SiteId int64, ) ([]map[string]string, error) `select Id,Title,UserId,date_format(ifnull(ReplyTime,CreateAt),'%Y-%m-%d %H:%i:%S') CreateAt from Topic where  SiteId=? order by Id desc limit 0,1000`

	//Top1000 func() ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt) CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime from Topic t left join Category c on t.CategoryId=c.Id
	//							  order by ifnull(t.ReplyTime,t.CreateAt) desc limit 0,1000`

	ListBySub func(SiteId int64, id, start interface{}) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt) CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime 
		from Topic t left join Category c on t.CategoryId=c.Id 
		where t.SiteId=? and CategoryId=? order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`

	ListByParent func(SiteId int64, parentId, start interface{}) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt) CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime
		from Topic t left join Category c on t.CategoryId=c.Id 
		where t.SiteId=? and CategoryId in (select Id from Category where ParentId=?) order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`

	ListByUserId func(SiteId int64, userId interface{}, start interface{}) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt) CreateAt,c.Name CategoryName,
	t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime,t.ShowTimes
	from Topic t left join Category c on t.CategoryId=c.Id
	where t.SiteId=? and t.UserId=? order by id desc limit ?,20`

	HotToday func(SiteId int64, ) ([]map[string]string, error) `select Id,UserId,Title,Today from Topic where SiteId=? Order by Today desc limit 0,10`
	HotAll   func(SiteId int64, ) ([]map[string]string, error) `select Id,UserId,Title,ShowTimes from Topic where SiteId=? Order by ShowTimes desc limit 0,10`
}

type DaoFav struct {
	List func(SiteId int64, UserId, start int64) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt)CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime
								from Topic t left join Category c on t.CategoryId=c.Id 
								where t.SiteId=? and t.Id in (select EntityId from Fav where UserId=?)
								order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`

	Count func(SiteId int64, UserId int64) (int64, bool, error)                 `select count(*) from Fav where SiteId=? and UserId=? `
	Exist func(SiteId int64, UserId int64, Id interface{}) (int64, bool, error) `select count(*) from Fav where SiteId=? and UserId=? and EntityId=?`
	Del   func(SiteId int64, UserId int64, Id interface{}) (int64, error)       `delete from Fav where SiteId=? and UserId=? and EntityId=?`
	Save  func(SiteId int64, UserId int64, Id interface{}) (int64, error)       `INSERT INTO Fav(SiteId,UserId,EntityId,CreateAt)VALUES(?,?,?,now())`
}
type DaoFollow struct {
	Exist    func(SiteId int64, Id int64, UserId interface{}) (int64, bool, error) `select count(*) from Follow where SiteId=? and FollowId=? and UserId=?`
	Follow   func(SiteId int64, FollowId, UserId int64) (int64, error)             `INSERT INTO Follow (SiteId,FollowId,UserId,CreateAt)VALUES(?,?,?,now())`
	UnFollow func(SiteId int64, FollowId, UserId int64) (int64, error)             `delete from Follow where SiteId=? and FollowId=? and UserId=?`
	Count    func(SiteId int64, UserId int64) (int64, bool, error)                 `select count(*) from Follow where SiteId=? and UserId=? `

	TopicList func(SiteId int64, UserId, start int64) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt)CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime
								from Topic t left join Category c on t.CategoryId=c.Id 
								where t.SiteId=? and t.UserId in (select FollowId from Follow where UserId=?)
								order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`
}

type DaoTopicAppend struct {
	List func(SiteId, TopicId int64) ([]map[string]string, error) `select Id,fmt(CreateAt)CreateAt,AppendText from TopicAppend where SiteId=? and TopicId=?`

	Count func(SiteId, TopicId int64) (int64, bool, error) `select count(*) from TopicAppend where SiteId=? and TopicId=?`

	Add func(SiteId, TopicId int64, AppendText string) (int64, error) `INSERT INTO  TopicAppend(SiteId,TopicId,AppendText,CreateAt)VALUES(?,?,?,now())`
}
