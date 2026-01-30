set -x
# should also work with
# GOOSPKG="github.com/abarisani/golang-proposal-73608@v0.0.2"

# user/linux amd64 example
GOOS=tamago GOARCH=amd64 GOOSPKG=${PWD} $TAMAGO run -ldflags '-X runtime.testBinary=true' main.go

# imx8mp evk arm64 example (WiP)
#GOOS=tamago GOARCH=arm64 GOOSPKG=${PWD} $TAMAGO run -ldflags "-T 0x40010000 -R 0x1000" main.go
#qemu-system-aarch64 -machine imx8mp-evk -m 512M -smp 1 -nographic -monitor none -semihosting -serial stdio -net nic,model=imx.enet,netdev=net0 -netdev tap,id=net0,ifname=tap0,script=no,downscript=no imx8mp_evk_arm64
