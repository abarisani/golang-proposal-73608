set -x
# should also work with
# GOOSPKG="github.com/abarisani/golang-proposal-73608@v0.0.2"
GOOS=tamago GOOSPKG=${PWD} $TAMAGO run -ldflags '-X runtime.testBinary=true' main_amd64.go
