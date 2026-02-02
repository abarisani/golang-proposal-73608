package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello from %s • %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
