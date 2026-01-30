package main

import (
	"fmt"

	// FIXME: for now goos imports must be loaded to cache any nostd module
	// imported by GOOSPKG.
	_ "github.com/usbarmory/tamago/arm64"
	_ "github.com/usbarmory/tamago/board/nxp/imx8mpevk"
	_ "github.com/usbarmory/tamago/soc/nxp/imx8mp"
)

func main() {
	fmt.Println("hello world")
}
