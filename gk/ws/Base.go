package ws

import (
	"github.com/ecdiy/gpa"
)

const PageSize = 20

var (
	Gpa *gpa.Gpa

	MultiSite = 0
)

func InitDao() {
	Gpa = gpa.Init(EnvParam("DbDriver"), EnvParam("DbDsn"), UserDao, KvDao, ScoreLog, MsgDao, TokenDao,

		TopicDao, TopicCategoryDao, ReplyDao, FavDao, FollowDao, AppendDao,

		NoteDao, CategoryDao,

		ZxDao, ThankDao, RefererDao,
	)
}
