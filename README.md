# CodeChain RPC Go

CodeChain RPC Go is a Go library that calls RPC to a CodeChain node.

## How to install

```bash
go get -u github.com/codechain-io/codechain-rpc-go
```

## Example

```Go
package main

import (
	"fmt"
	rpc "github.com/codechain-io/codechain-rpc-go"
)
func main() {
	fmt.Println(codechainrpc.Ping(nil))
}

```
