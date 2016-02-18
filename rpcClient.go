package jsonrpc

import (
	"jsonrpc/transport"
	"strings"
)

const (
	JSON_RPC_ID int8 = 0
)

type RpcClient struct {
	Method   string
	Urls     transport.Url
	Field    map[string]interface{}
	Header   []string
	Code     int16
	Error    error
	Response interface{}
}

// set
func (client *RpcClient) setMethod(method string) {
	client.Method = strings.ToUpper(method)
}

func (client *RpcClient) setUrls(urls transport.Url) {
	client.Urls = urls
}

func (client *RpcClient) setField(field map[string]interface{}) {
	client.Field = field
}

func (client *RpcClient) setHeader(header []string) {
	client.Header = header
}

// get todo

// call to get response
func (client *RpcClient) call(method string, params map[string]interface{}) {
	client.setField(map[string]interface{}{
		"id": JSON_RPC_ID + 1,
		"method": method,
		"params": params,
	})
	if len(client.Method) == 0 {
		client.setMethod("POST")
	}
	if len(client.Header) == 0 {
		client.setHeader(make([]string, 2, 2))
	}
	// todo
	_, _, _, _ = transport.Send(client.Method, client.Urls, client.Field, client.Header)
}