# Pokémon Go Protobuffers for Golang
Exported from and maintained at https://github.com/AeonLucid/POGOProtos

## Usage

```go
package main

import (
  "fmt"
  p "github.com/pkmngo-odi/pogo-protos"
)

func main() {
  fmt.Println(p.TeamColor_BLUE)
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
$ python ./compile_single.py -l=go --out_path=$GOPATH/src/github.com/pkmngo-odi/pogo-protos --go_root_package=github.com/pkmngo-odi/pogo-protos
```
