package ws

import "encoding/json"

type ScoreRule struct {
	Thank, LoginMax, LoginMin, Reply, GetReply, Topic, Register int64
}

var ScoreMap = make(map[int64]*ScoreRule)

func GetSoreRule(siteId int64) *ScoreRule {
	r, b := ScoreMap[siteId]
	if b {
		return r
	}
	rule, _, _ := KvDao.Get(siteId, "ScoreRule")
	return SetSoreRule(siteId, rule)
}
func SetSoreRule(siteId int64, rule string) *ScoreRule {
	if len(rule) > 2 {
		df := &ScoreRule{}
		je := json.Unmarshal([]byte( rule), &df)
		if je == nil {
			ScoreMap[siteId] = df
			return df
		}
	}
	dfx := &ScoreRule{Thank: -10, LoginMax: 50, LoginMin: 5, Reply: -5, GetReply: 5, Topic: -20, Register: 2000}
	ScoreMap[siteId] = dfx
	return dfx
}
