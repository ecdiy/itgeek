package ws

var (
	ZxDao = &DaoZx{}
)

type DaoZx struct {
	Add func(SiteId int64, PageKey, Ua, TestJson string) (int64, error) `INSERT INTO Zx(SiteId,PageKey,Ua,TestJson,CreateAt)VALUES(?,?,?,?,now())`

	UpJson func(TestJson, Ua, PageKey string, SiteId int64) (int64, error) `update Zx set TestJson=? where Ua=? and PageKey=? and SiteId=?`

	Pub func(Ua, PageKey string, SiteId int64) (int64, error) `update Zx set ZxJson=TestJson where Ua=? and PageKey=? and SiteId=?`

	GetTest func(SiteId int64, Ua, PageKey string) (string, bool, error) `select TestJson from Zx where SiteId=? and Ua=? and PageKey=?`
	Get     func(SiteId int64, Ua, PageKey string) (string, bool, error) `select ZxJson from Zx where SiteId=? and Ua=? and PageKey=?`
}
