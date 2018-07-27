package ws

import (
	"github.com/ecdiy/gpa"
	"github.com/gin-gonic/gin"
	"strings"
)

const PageSize = 20

var (
	Gpa       *gpa.Gpa
	TokenMap  = make(map[string]map[string]string)
	MultiSite = 0

	ScoreMap = make(map[int64]*ScoreRule)
	WebGin   = gin.New()
)

func InitDao() {
	Gpa = gpa.Init(EnvParam("DbDriver"), EnvParam("DbDsn"), UserDao, KvDao, ScoreLog, MsgDao, TokenDao,

		TopicDao, TopicCategoryDao, ReplyDao, FavDao, FollowDao, AppendDao,

		NoteDao, CategoryDao,

		ZxDao, ThankDao, RefererDao,
	)
}

func SiteId(c *gin.Context) int64 {
	if MultiSite == 1 {
		sql := "select SiteId from site.BindDomain where Domain=?"
		host := c.Request.Host
		ix := strings.Index(host, ":")
		if ix > 0 {
			host = host[0:ix]
		}
		idx := strings.Index(host, ".ecdiy.cn")
		if idx > 0 {
			sql = "select Id from site.Site where SubDomain=?"
			host = host[0:idx]
		}
		hId, qb, _ := Gpa.QueryInt64(sql, host)
		if qb {
			return hId
		}
		return 0
	}
	return 0
}

func WebAuth(url string, fun func(wdb *Web)) {
	WebGin.POST(url, func(c *gin.Context) {
		web := Verify(c)
		if web.Auth {
			web.initParam()
			fun(web)
			c.JSON(200, web.Out)
		} else {
			c.AbortWithStatus(401)
		}
	})
}

func WebPost(url string, fun func(wdb *Web)) {
	WebGin.POST(url, func(c *gin.Context) {
		web := Verify(c)
		fun(web)
		c.JSON(200, web.Out)
	})
}
