package jsonrpc

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/rpc/v2/json2"
)

const (
	rpcURL  string = "138.201.255.176"
	rpcPort int    = 8545
)

// ClientRequest - 送出基於 JSON RPC 2.0 格式的請求
func ClientRequest(method string, params interface{}) *simplejson.Json {
	data, _ := json2.EncodeClientRequest(method, params)
	request, _ := http.NewRequest("POST", fmt.Sprintf("http://%s:%d", rpcURL, rpcPort), bytes.NewBuffer(data))

	request.Header.Set("Content-Type", "application/json")
	reply, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer reply.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reply.Body)
	response, err := simplejson.NewJson(buffer.Bytes())
	if err != nil {
		fmt.Println("parse JSON fail")
	}

	return response
}
