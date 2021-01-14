package test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2021/1/11 下午4:13
 * @company    ：eeo.cn
 */

func TestCurl(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(`recived err : `, err)
		}
		t.Log(`业务走不下去了，来这里`)
	}()
	ConsulReq()
	t.Log(`正常的业务啊啊啊啊啊`)
}

func ConsulReq() []byte {
	url := "http://127.0.0.1:8500/v1/agent/members"
	payload := strings.NewReader("{\"Name\":\"dxx\",\"Type\":\"manageement\"}")
	req, _ := http.NewRequest("GET", url, payload)
	req.Header.Add("x-consul-token", "p2BE1AtpwPbrxZdC6k+eXA==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "9c10eb08-08a3-ee2d-7160-570f851b5cab")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
