package ws

import (
	"github.com/ecdiy/gpa"
)

var (
	Gpa *gpa.Gpa
)

func InitDao() {
	Gpa = gpa.Init(EnvParam("Dsn"), EnvParam("Dsn"), UserDao, KvDao, ScoreLog, MsgDao, TokenDao,

		TopicDao, TopicCategoryDao, ReplyDao, FavDao, FollowDao,

		NoteDao, CategoryDao,

		ZxDao, ThankDao, RefererDao,
	)
}
