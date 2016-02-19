package transport
import (
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/pquerna/ffjson/ffjson"
)

func parseGetData(request *http.Request) interface{} {
	var json interface{}
	ffjson.Unmarshal([]byte(request.Form["content"][0]), &json)
	return json
}

func parsePostData(request *http.Request) interface{} {
	result, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	var json interface{}
	ffjson.Unmarshal(result, &json)
	m := json.(map[string]interface{})
	return m
}

func ParseRequest(request *http.Request) interface{} {
	var data interface{}
	//解析参数
	request.ParseForm()
	if method := strings.ToUpper(request.Method); method == "GET" {
		data = parseGetData(request)
	} else if method == "POST" {
		data = parsePostData(request)
	}
	return data
}