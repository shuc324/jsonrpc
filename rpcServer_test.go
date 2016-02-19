package jsonrpc

import (
	"net/http"
	"log"
	"testing"
	"jsonrpc/transport"
)

func Test_Run(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(transport.ParseRequest(r))
}
