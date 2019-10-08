package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Result of RPC
type Result struct {
	Jsonrpc string          `json:"jsonrpc,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	ID      interface{}     `json:"id,omitempty"`
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
	if idGen == nil {
		i := 0
		idGen = func() string {
			i++
			return strconv.Itoa(i)
		}
	}
	return rpcClient{nodeURL, idGen, client}
}

func (client *rpcClient) id(option callInterface) interface{} {
	if client.idGenerator != nil {
		return client.idGenerator()
	}
	return nil
}

func (client *rpcClient) call(option callInterface, result interface{}, params ...interface{}) {
	defer func() {
		s := recover()
		if s != nil {
			fmt.Println(s)
		}
	}()

	reqBodyMap := map[string]interface{}{"jsonrpc": "2.0", "method": option.method, "params": params, "id": client.idGenerator()}
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
	var outerResult Result
	if err := json.Unmarshal(body, &outerResult); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(outerResult.Result, &result); err != nil {
		panic(err)
	}

	// TODO: handle errors
}
