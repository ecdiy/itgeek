package ws

var ResourceDao = &DaoResource{}
var ResourceSummaryDao = &DaoResourceSummary{}

type DaoResource struct {
	Add func(SiteId, UserId, Size int64, FileName string) (int64, error) `INSERT INTO Resource(SiteId,UserId,Size,FileName,CreateAt,Status)VALUES(?,?,?,?,now(),0)`

	Up func(ResName, ResDesc string, ResScore, Id, userId int64) (int64, error) `update Resource set ResName=?,ResDesc=?,ResScore=? where Id=? and UserId=?`
}

type DaoResourceSummary struct {
	Get func(SiteId, UserId int64) (map[string]string, bool, error) `select * from ResourceSummary where SiteId=? and UserId=?`
}
