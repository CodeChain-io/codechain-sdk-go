# CodeChain SDK Go

A Golang SDK for CodeChain

## How to install

```bash
go get -u github.com/CodeChain-io/codechain-sdk-go
```

## Example

**Check for** [examples folder](https://github.com/CodeChain-io/codechain-sdk-go/tree/master/example)

Examples works under CodeChain in local with [test configuration](https://github.com/CodeChain-io/codechain#run).

```sh
go run ./example/pay/main.go
```

```Go
package main

import (
	"fmt"
	"github.com/CodeChain-io/codechain-sdk-go/rpc"
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
