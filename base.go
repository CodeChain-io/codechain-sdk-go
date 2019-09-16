package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Result of RPC
type Result struct {
	Jsonrpc string      `json:"jsonrpc,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	ID      string      `json:"id,omitempty"`
}

type callInterface struct {
	method string
	id     string
}

type rpcClient struct {
	rpcURL      string
	idGenerator func() string
	httpClient  http.Client
}

func newRPCClient(nodeURL string, idGen func() string) rpcClient {
	client := http.Client{}
	return rpcClient{nodeURL, idGen, client}
}

func (client *rpcClient) id(option callInterface) interface{} {
	if client.idGenerator != nil {
		return client.idGenerator()
	}
	return nil
}

func (client *rpcClient) call(option callInterface, params ...interface{}) Result {

	defer func() {
		s := recover()
		if s != nil {
			fmt.Println(s)
		}
	}()

	reqBodyMap := map[string]interface{}{"jsonrpc": "2.0", "method": option.method, "params": params, "id": ""}
	reqBodyJSON, _ := json.Marshal(reqBodyMap)
	reqBody := bytes.NewBuffer(reqBodyJSON)
	req, err := http.NewRequest("POST", client.rpcURL, reqBody)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, readerr := ioutil.ReadAll(resp.Body)
	if readerr != nil {
		panic(readerr)
	}
	var res Result
	json.Unmarshal(body, &res)
	return res

	// TODO: handle errors
}
