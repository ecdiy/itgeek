package ws

var (
	NoteDao     = &DaoNote{}
	CategoryDao = &DaoCategory{}
)

type DaoCategory struct {
	Add func(SiteId int64, Name string, ParentId, ItemCount, UserId int64) (int64, error) `INSERT INTO NoteCategory(SiteId,Name,ParentId,CreateAt,ModifyAt,ItemCount,UserId)VALUES (?,?,?,now(),now(),?, ?)`

	List        func(SiteId int64, userId int64) ([]map[string]string, error)    `select Id,Name,ParentId from NoteCategory where SiteId=? and UserId=?`
	ModifyName  func(Name string, SiteId int64, Id, UserId int64) (int64, error) `update NoteCategory set Name=? where SiteId=? and Id=? and UserId=?`
	Del         func(SiteId int64, Id, UserId int64) (int64, error)              `Delete from NoteCategory where SiteId=? and  Id=? and UserId=?`
	DelParentId func(SiteId int64, Id, UserId int64) (int64, error)              `Delete from NoteCategory where SiteId=? and  ParentId=? and UserId=?`
}

type DaoNote struct {
	Detail func(SiteId int64, Id, UserId int64) (map[string]string, bool, error) `select Title,Body from Notes where SiteId=? and Id=? and UserId=?`

	Add func(SiteId int64, NoteCategoryId, Title, Body, SourceType, Source string, UserId int64) (int64, error) `INSERT INTO Notes(SiteId,CategoryId,Title,Body,CreateAt,SourceType,Source,UserId)VALUES(?,?,?,?,now(),?,?,?)`

	List       func(SiteId int64, userId int64) ([]map[string]string, error) `select Id,Title,CategoryId from Notes where SiteId=? and UserId=?`
	UpdateCat  func(catId, CatId, UserId int64, SiteId int64) (int64, error) `update Notes set CategoryId=? where CategoryId=? and UserId=? and SiteId=?`
	UpdatePCat func(catId, CatId, UserId int64, SiteId int64) (int64, error) `update Notes set CategoryId=0 where CategoryId in(select Id from NoteCategory where ParentId=?) and UserId=?  and SiteId=?`
}
