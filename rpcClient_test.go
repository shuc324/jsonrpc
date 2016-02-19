package jsonrpc

import (
	"testing"
	"log"
	"jsonrpc/transport"
)

func Test_Rpc_Client_Run(t *testing.T)  {
	rpcClient := RpcClient{}

	rpcClient.setMethod("get").setUrls(transport.Url{"http://127.0.0.1", 9001, ""}).setUri("common").call("login", []interface{}{"name", "shuc", "age", 21});
	rpcClient.setMethod("post").setUrls(transport.Url{"http://127.0.0.1", 9001, ""}).setUri("common").call("login", []interface{}{"name", "shuc", "age", 21});

	log.Println(rpcClient.getResponse())
}
