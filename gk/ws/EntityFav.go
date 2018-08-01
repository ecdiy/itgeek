package ws

const (
	FavTypeTopic = 1
	FavTypeDown  = 2
)

var FavDao = &DaoFav{}

type DaoFav struct {
	ListFavTopic func(SiteId int64, UserId, start int64) ([]map[string]string, error) `select t.Id,t.Title,t.Username,t.UserId,fmt(t.CreateAt)CreateAt,c.Name CategoryName,t.CategoryId,t.ReplyCount,t.ReplyUsername,t.ReplyUserId,ifnull(t.ReplyTime,t.CreateAt)ReplyTime
								from Topic t left join Category c on t.CategoryId=c.Id 
								where t.SiteId=? and t.Id in (select EntityId from Fav where UserId=? and FavType=1)
								order by ifnull(t.ReplyTime,t.CreateAt) desc limit ?,20`

	Count func(SiteId int64, UserId int64) (int64, bool, error) `select count(*) from Fav where SiteId=? and UserId=? `

	Exist func(SiteId int64, UserId int64, Id interface{}) (int64, bool, error) `select count(*) from Fav where SiteId=? and UserId=? and EntityId=?`

	Del func(SiteId int64, UserId int64, Id interface{}) (int64, error) `delete from Fav where SiteId=? and UserId=? and EntityId=?`

	Save func(SiteId int64, UserId, FavType int64, Id interface{}) (int64, error) `INSERT INTO Fav(SiteId,UserId,FavType,EntityId,CreateAt)VALUES(?,?,?,?,now())`
}
