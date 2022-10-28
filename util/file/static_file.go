package file

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func StaticFile(u *gin.Context) {
	url := u.Query("url")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		u.JSON(http.StatusOK, errors.New("client.Do(req)"+err.Error()))
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	imgBytes, _ := ioutil.ReadAll(resp.Body)

	u.Data(http.StatusOK, "utf-8", imgBytes)
}
