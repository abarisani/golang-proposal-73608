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

import "unsafe"

// defined in syscall_*.s
func CPUInit()
func sys_exit(code int32)
func sys_write(c *byte)
func sys_clock_gettime() (ns int64)
func sys_getrandom(b []byte, n int)

var (
	RamStart       uint = 0x80000000
	RamSize        uint = 0x20000000
	RamStackOffset uint = 0x100

	Bloc   = uintptr(RamStart)
	Exit   = sys_exit
	Idle   func(until int64)
	ProcID func() uint64
	Task   func(sp, mp, gp, fn unsafe.Pointer)

	Hwinit0  = func() {}
	InitRNG  = func() {}
	Nanotime = sys_clock_gettime
	Hwinit1  = func() {}
)

func GetRandomData(b []byte) {
	sys_getrandom(b, len(b))
}

// preallocated memory to avoid malloc during panic
var a [1]byte

func Printk(c byte) {
	a[0] = c
	sys_write(&a[0])
}
