# jsonrpc

Client:

    rpcClient := RpcClient{}

    rpcClient.setMethod("get").setUrls(transport.Url{"http://www.baidu.com", 80, ""}).setUri("common").call("login", map[string]interface{}{"name": "shuc", "age": 21});

Server:

    todo