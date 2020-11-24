package cons

import (
	"github.com/dreamlu/gt/tool/conf"
)

// "智能推荐", "玄幻", "仙侠修真", "科幻", "网游游戏", "历史穿越", "都市"
const (
	IntelRec  = "推荐"
	Classify1 = "玄幻"
	Classify2 = "仙侠修真"
	Classify3 = "科幻"
	Classify4 = "网游游戏"
	Classify5 = "历史穿越"
	Classify6 = "都市"
)

var (
	DevMode = conf.GetString("app.devMode")
	Version = conf.GetString("app.version")
)
