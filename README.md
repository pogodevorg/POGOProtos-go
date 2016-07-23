# Pokémon Go Protobuffers for Golang
Exported from and maintained at https://github.com/AeonLucid/POGOProtos

## Usage

```go
package main

import (
  "fmt"
  "github.com/zeeraw/pogo-protos/data"
)

func main() {
  fmt.Println(data.TeamColor_BLUE)
}
```

## Update this repository
Install the Go protoc extensions through `go get`

```bash
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

Check out the [AeonLucid/POGOProtos](https://github.com/AeonLucid/POGOProtos) repository and run their scripts.

```bash
$ git clone git@github.com:AeonLucid/POGOProtos.git
$ cd POGOProtos
$ python ./compile_single.py -l=go --out_path=$GOPATH/src/github.com/zeeraw/pogo-protos --go_root_package=github.com/zeeraw/pogo-protos
```
