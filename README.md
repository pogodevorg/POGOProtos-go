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

Run the update script which will commit any changes.

```bash
$ ./update.sh
```
