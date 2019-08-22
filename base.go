package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	node        string
	idGenerator func(string) string
	client      http.Client
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

func newRPCClient(nodeURL string, idGen func(string) string) {
	node = nodeURL
	idGenerator = idGen
	client = http.Client{}
}

func id(option callInterface) interface{} {
	if idGenerator != nil {
		return idGenerator(option.method)
	}
	return nil
}

func call(option callInterface, params ...interface{}) Result {

	defer func() {
		s := recover()
		if s != nil {
			fmt.Println(s)
		}
	}()

	reqBodyMap := map[string]interface{}{"jsonrpc": "2.0", "method": option.method, "params": params, "id": ""}
	reqBodyJSON, _ := json.Marshal(reqBodyMap)
	reqBody := bytes.NewBuffer(reqBodyJSON)
	req, err := http.NewRequest("POST", node, reqBody)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
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
