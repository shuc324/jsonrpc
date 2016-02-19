package jsonrpc

import (
	"testing"
	"log"
	"jsonrpc/transport"
)

func Test_Rpc_Client_Run(t *testing.T)  {
	rpcClient := RpcClient{}

	rpcClient.setMethod("get").setUrls(transport.Url{"http://www.baidu.com", 80, ""})


	log.Println(rpcClient)

}
