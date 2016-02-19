# jsonrpc

Client:

    rpcClient := RpcClient{}

    rpcClient.setMethod("post").setUrls(transport.Url{"http://127.0.0.1", 9001, ""}).setUri("common").call("login", []interface{}{"name", "shuc", "age", 21});

Server:

    todo