package transport
import (
	"testing"
	"log"
)

// 测试client
func Test_Client_Run(t *testing.T) {
	url := Url{}
	url.Server = "http://www.baidu.com"
	url.Port = 80
	url.Uri = ""

	filed := make(map[string]interface{}, 10)
	header := make([]string, 1)
	res, code, _, response := Send("POST", url, filed, header)
	if code != 0 {
		log.Fatal(response)
	}
	log.Println(res)
}
