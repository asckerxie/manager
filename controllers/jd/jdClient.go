package jd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var url_access_token = "https://kploauth.jd.com/oauth/token"

	var paraMap map[string]string

	// 再使用make函数创建一个非nil的map，nil map不能赋值
	paraMap = make(map[string]string)

	paraMap["grant_type"] = "password"
	paraMap["app_key"] = "45ac35bb34ef46dbb8fa988059eb3dec"
	paraMap["app_secret"] = "9e0269931118418aa4133a3c0bc119c8"
	paraMap["state"] = "0"
	paraMap["password"] = "e10adc3949ba59abbe56e057f20f883e"
	paraMap["username"] = "%e5%a5%b6%e7%93%b6%e7%90%86%e8%b4%a2"

	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", url_access_token, nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}

}
