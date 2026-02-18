# user/linux amd64 example, using GOROOT/src/runtime/goos as GOOSPKG
GOOS=tamago GOARCH=amd64 $TAMAGO run main_amd64.go

# imx8mp evk arm64 example
GOOS=tamago GOARCH=arm64 GOOSPKG=github.com/usbarmory/tamago $TAMAGO build -ldflags "-T 0x40010000 -R 0x1000" main_imx8mpevk.go && \
echo "launching QEMU (Ctrl+C to quit)" && \
qemu-system-aarch64 \
	-machine imx8mp-evk -m 6G -smp 1 \
	-nographic -monitor none -semihosting -serial stdio \
	-kernel main_imx8mpevk
