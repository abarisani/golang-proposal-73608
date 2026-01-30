
This repository is a PoC for the following Go proposal:

  * [proposal: all: add bare metal support #73608](https://github.com/golang/go/issues/73608)

It requires the following modified Go distribution:

  * https://github.com/abarisani/tamago-go/tree/tamago1.26rc2-73608

Example:

```
GOOS=tamago GOARCH=amd64 GOOSPKG="github.com/abarisani/golang-proposal-73608@v0.0.3" go run -ldflags '-X runtime.testBinary=true' main.go
go: downloading github.com/abarisani/golang-proposal-73608 v0.0.3
hello world
```
