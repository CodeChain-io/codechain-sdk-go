# CodeChain SDK Go

A Golang SDK for CodeChain

## How to install

```bash
go get -u github.com/CodeChain-io/codechain-rpc-go
```

## Example

```Go
package main

import (
	"fmt"
	rpc "github.com/CodeChain-io/codechain-rpc-go"
)
func main() {
	a := rpc.NewRPC("https://corgi-rpc.codechain.io/").Chain
	b, _ := a.GetBlockByNumber(0)

	fmt.Printf("%+v\n", b)
}

```

## How to run linter

Make sure you run `revive` before creating a PR to the repo

### Install

```sh
go get -u github.com/mgechev/revive
```

### Run

```sh
revive -config revive.toml ./...
```
