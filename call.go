package rpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type callInterface struct {
	node   string      //`json:"node"`
	method string      //`json:"method"`
	id     interface{} //`json:"id"`
}

func call(option callInterface, params ...[]interface{}) string {

	reqBodyMap := map[string]interface{}{"jsonrpc": "2.0", "method": option.method, "params": params, "id": option.id}
	reqBodyJSON, _ := json.Marshal(reqBodyMap)
	reqBody := bytes.NewBuffer(reqBodyJSON)

	req, err := http.NewRequest("POST", option.node, reqBody)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, readerr := ioutil.ReadAll(resp.Body)
	if readerr != nil {
		panic(readerr)
	} else {
		return string(body)
	}
	// TODO: handle errors
}
