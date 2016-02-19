package jsonrpc

import (
	"jsonrpc/transport"
	"strings"
)

const (
	JSON_RPC_ID int8 = 0
)

// rpc client
type RpcClient struct {
	Method   string
	Urls     transport.Url
	Field    map[string]interface{}
	Header   []string
	Code     int16
	Error    error
	Response interface{}
	Body     string
}

// set
func (client *RpcClient) setMethod(method string) *RpcClient {
	client.Method = strings.ToUpper(method)
	return client
}

func (client *RpcClient) setUri(uri string) *RpcClient {
	client.Urls.Uri = uri
	return client
}

func (client *RpcClient) setUrls(urls transport.Url) *RpcClient {
	client.Urls = urls
	return client
}

func (client *RpcClient) setField(field map[string]interface{}) *RpcClient {
	client.Field = field
	return client
}

func (client *RpcClient) setHeader(header []string) *RpcClient {
	client.Header = header
	return client
}

func (client *RpcClient) setCode(code int16) *RpcClient {
	client.Code = code
	return client
}

func (client *RpcClient) setError(err error) *RpcClient {
	client.Error = err
	return client
}

func (client *RpcClient) setResponse(response interface{}) *RpcClient {
	client.Response = response
	return client
}

func (client *RpcClient) setBody(body string) *RpcClient {
	client.Body = body
	return client
}

// get
func (client *RpcClient) getCode() (code int16) {
	return client.Code
}

func (client *RpcClient) getResponse() (response interface{}) {
	return client.Response
}

func (client *RpcClient) getError() (err error) {
	return client.Error
}

func (client *RpcClient) getBody() (body string) {
	return client.Body
}

func (client *RpcClient) call(method string, params map[string]interface{}) *RpcClient {
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
	client.Response, client.Code, client.Error, client.Body = transport.Send(client.Method, client.Urls, client.Field, client.Header)
	return client
}
