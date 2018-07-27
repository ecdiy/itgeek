package ws

import (
	"github.com/ecdiy/gpa"
	"github.com/gin-gonic/gin"
	"strings"
	"regexp"
)

const PageSize = 20

var (
	Gpa       *gpa.Gpa
	TokenMap  = make(map[string][]string)
	MultiSite = 0

	ScoreMap = make(map[int64]*ScoreRule)
	WebGin   = gin.New()

	UaSeo = regexp.MustCompile(`baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator`)
	UaH5  = regexp.MustCompile(`(MIDP)|(WAP)|(UP.Browser)|(Smartphone)|(Obigo)|(Mobile)|(AU.Browser)|(wxd.Mms)|(WxdB.Browser)|(CLDC)|(UP.Link)|(KM.Browser)|(UCWEB)|(SEMC\-Browser)|(Mini)|(Symbian)|(Palm)|(Nokia)|(Panasonic)|(MOT\-)|(SonyEricsson)|(NEC\-)|(Alcatel)|(Ericsson)|(BENQ)|(BenQ)|(Amoisonic)|(Amoi\-)|(Capitel)|(PHILIPS)|(SAMSUNG)|(Lenovo)|(Mitsu)|(Motorola)|(SHARP)|(WAPPER)|(LG\-)|(LG/)|(EG900)|(CECT)|(Compal)|(kejian)|(Bird)|(BIRD)|(G900/V1.0)|(Arima)|(CTL)|(TDG)|(Daxian)|(DAXIAN)|(DBTEL)|(Eastcom)|(EASTCOM)|(PANTECH)|(Dopod)|(Haier)|(HAIER)|(KONKA)|(KEJIAN)|(LENOVO)|(Soutec)|(SOUTEC)|(SAGEM)|(SEC\-)|(SED\-)|(EMOL\-)|(INNO55)|(ZTE)|(iPhone)|(Android)|(Windows CE)|(Wget)|(Java)|(curl)|(Opera)/`)
)

func InitWs() {

	MultiSite = EnvParamInt("MultiSite", 0)
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
		web.initParam()
		fun(web)
		c.JSON(200, web.Out)
	})
}

func GetUa(ctx *gin.Context) string {
	ua := ctx.Request.UserAgent()
	if len(ua) == 0 {
		return "web"
	}
	if UaH5.MatchString(ua) {
		return "h5"
	}
	if UaSeo.MatchString(ua) {
		return "seo"
	}
	return "web"
}
