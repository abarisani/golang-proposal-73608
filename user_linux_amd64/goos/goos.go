// Linux user space support
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package goos provides support for using `GOOS=tamago` in Linux user
// space.
//
// This package is only meant to be used with `GOOS=tamago` as supported by the
// TamaGo framework for bare metal Go, see https://github.com/usbarmory/tamago.
package goos

import (
	"unsafe"
)

var (
	RamStart       uint64 = 0x80000000
	RamSize        uint64 = 0x20000000
	RamStackOffset uint64 = 0x100
	Exit                  = sys_exit

	Idle   func(code int64)                    = nil
	ProcID func() uint64                       = nil
	Task   func(sp, mp, gp, fn unsafe.Pointer) = nil
)

// defined in syscall_*.s
func sys_exit(code int32)
func sys_write(c *byte)
func sys_clock_gettime() (ns int64)
func sys_getrandom(b []byte, n int)

func Hwinit0() {}
func Hwinit1() {}
func InitRNG() {}

func Nanotime() int64 {
	return sys_clock_gettime()
}

// preallocated memory to avoid malloc during panic
var a [1]byte

func Printk(c byte) {
	a[0] = c
	sys_write(&a[0])
}

func GetRandomData(b []byte) {
	sys_getrandom(b, len(b))
}

func Bloc() uintptr {
	return uintptr(RamStart)
}
