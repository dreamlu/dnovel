package request

import (
	"errors"
	"github.com/dreamlu/gt/src/ghttp"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func StaticFile(u *gin.Context) {
	url := u.Query("url")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		u.JSON(http.StatusOK, errors.New("client.Do(req)"+err.Error()))
		return
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	imgBytes, _ := ioutil.ReadAll(resp.Body)

	u.Data(http.StatusOK, "utf-8", imgBytes)
}

func RS(u *gin.Context) {
	urlT := u.Query("url")
	method := u.Query("method")
	values, _ := url.Parse(urlT)
	rs := strings.Split(values.RawQuery, "=")

	r := ghttp.NewRequest(method, urlT)
	l := len(rs)
	for i := 0; i < l && l > 1; i += 2 {
		r.SetForm(rs[i], rs[i+1])
	}
	r.SetContentType(ghttp.ContentTypeFormUrl)
	r.SetHeader("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	r.SetHeader("cookie", "cf_clearance=6f520d523c83408ee2ccbb81fab2588ac9b8352a-1667284728-0-150")
	res := r.Exec()
	if res.Error() != nil {
		u.JSON(http.StatusOK, errors.New("client.Do(req)"+res.Error().Error()))
		return
	}
	u.Data(http.StatusOK, "content-type: text/html; charset=utf-8", res.MustBytes())
}
