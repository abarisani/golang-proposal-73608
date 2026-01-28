set -x
GOOS=tamago GOOSPKG="${PWD}/user_linux_amd64" $TAMAGO run -ldflags '-X runtime.testBinary=true' main_amd64.go
