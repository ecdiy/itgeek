package ws

import (
	"github.com/ecdiy/gpa"
	"github.com/gin-gonic/gin"
)

const PageSize = 20

var (
	Gpa       *gpa.Gpa
	TokenMap  = make(map[string]map[string]string)
	MultiSite = 0

	WebGin = gin.New()
)

func InitDao() {
	Gpa = gpa.Init(EnvParam("DbDriver"), EnvParam("DbDsn"), UserDao, KvDao, ScoreLog, MsgDao, TokenDao,

		TopicDao, TopicCategoryDao, ReplyDao, FavDao, FollowDao, AppendDao,

		NoteDao, CategoryDao,

		ZxDao, ThankDao, RefererDao,
	)
}
