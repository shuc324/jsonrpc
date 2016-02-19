package transport

import (
	"net/http"
	"strings"
	"github.com/pquerna/ffjson/ffjson"
	"bytes"
	"io/ioutil"
	"net/url"
	"strconv"
)

const (
	SUCCESS, JSON_MARSHAL_ERROR, JSON_UN_MARSHAL_ERROR, IO_READ_ERROR, URL_PARSE_ERROR, HTTP_GET_ERROR, HTTP_POST_ERROR int16 = 0, -100, -200, -300, -400, -500, -600
)

type Url struct {
	Server string
	Port   int
	Uri    string
}

func (url *Url) toString() string {
	return url.Server + ":" + strconv.Itoa(url.Port) + "/" + url.Uri
}

type JsonRpcResult interface{}

// 字符串数组去除重复和空
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sLen := len(a)
	for i := 0; i < sLen; i++ {
		if (i > 0 && a[i - 1] == a[i]) || len(a[i]) == 0 {
			continue;
		}
		ret = append(ret, a[i])
	}
	return
}

// 发送请求
func Send(method string, urls Url, field map[string]interface{}, header []string) (res interface{}, code int16 , err error, body string) {
	var response string
	var jsonRpcRes JsonRpcResult
	b, err := ffjson.Marshal(field)
	if err != nil {
		return nil, JSON_MARSHAL_ERROR, err, ""
	}
	if method = strings.ToUpper(method); method == "GET" {
		u, err := url.Parse(urls.toString())
		if err != nil {
			return nil, URL_PARSE_ERROR, err, ""
		}
		q := u.Query()
		q.Set("content", string(b))
		u.RawQuery = q.Encode()
		res, err := http.Get(u.String());
		if err != nil {
			return nil, HTTP_GET_ERROR, err, ""
		}
		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, IO_READ_ERROR, err, ""
		}
		response := string(result)
		err = ffjson.Unmarshal(result, jsonRpcRes)
		if err != nil {
			return nil, JSON_UN_MARSHAL_ERROR, err, response
		}
	} else {
		header = append(header, "application/json")
		header = append(header, "charset=utf-8")
		header = RemoveDuplicatesAndEmpty(header)
		headerString := ""
		for _, item := range header {
			headerString = headerString + ";" + item
		}
		body := bytes.NewBuffer([]byte(b))
		res, err := http.Post(urls.toString(), headerString, body)
		if err != nil {
			return nil, HTTP_POST_ERROR, err, ""
		}
		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return nil, IO_READ_ERROR, err, ""
		}
		response := string(result)
		err = ffjson.Unmarshal(result, jsonRpcRes)
		if err != nil {
			return nil, JSON_UN_MARSHAL_ERROR, err, response
		}
	}
	return jsonRpcRes, 0, nil, response
}