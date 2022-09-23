package novel

import (
	"strings"
)

// 来源biquge：
// 破解后的增加的url参数
func resBody(body []byte) string {
	res := string(body)
	if strings.Contains(res, "<head>") {
		return ""
	}

	ress := strings.Split(res, "searchkey=")
	res = strings.Split(ress[1], `";`)[0]
	ss := strings.Split(res, "&")
	res = "&" + ss[len(ss)-1]
	return res
}
