package ws

import (
	"github.com/ecdiy/gpa"
	"github.com/gin-gonic/gin"
	"regexp"
)

const (
	PageSize    = 20
	KvGeekAdmin = "GeekAdmin"
)

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
func KeySave(siteId int64, k, v string) (int64, error) {
	_, e, _ := KvDao.Get(siteId, k)
	if e {
		return KvDao.Update(k, v, siteId)
	} else {
		return KvDao.Add(siteId, k, v)
	}
}
