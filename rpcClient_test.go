package jsonrpc

import (
	"testing"
	"jsonrpc/transport"
	"log"
)

func Test_Rpc_Client_Run(t *testing.T)  {
	rpcClient := RpcClient{}

	rpcClient.setMethod("get").setUrls(transport.Url{"http://www.baidu.com", 80, ""}).setUri("common").call("login", map[string]interface{}{"name": "shuc", "age": 21});

	log.Println(rpcClient.getResponse())
}
