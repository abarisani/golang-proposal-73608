
This repository is a PoC for the following Go proposal:

  * [proposal: all: add bare metal support #73608](https://github.com/golang/go/issues/73608)

It requires the following modified Go distribution:

  * https://github.com/abarisani/tamago-go/tree/tamago1.26rc2-73608

Example:

```
GOOS=tamago GOARCH=amd64 go run main.go
hello world

GOOS=tamago GOARCH=arm64 GOOSPKG="github.com/usbarmory/tamago@v0.0.0-20260202162439-a56f03727d4a" go build -ldflags "-T 0x40010000 -R 0x1000" main.go
qemu-system-aarch64 \
	-machine imx8mp-evk -m 6G -smp 1 \
	-nographic -monitor none -semihosting -serial stdio \
	-net nic,model=imx.enet,netdev=net0 -netdev tap,id=net0,ifname=tap0,script=no,downscript=no \
	-kernel main_imx8mpevk
```

For a complete port of [tamago-example](https://github.com/usbarmory/tamago-example) to proposal #73608 see:

  * https://github.com/abarisani/tamago-example/tree/go1.26-goos
