// 8MPLUSLPD4-EVK support for tamago/arm64
// https://github.com/usbarmory/tamago
//
// Copyright (c) The TamaGo Authors. All Rights Reserved.
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

// Package goos provides support for using `GOOS=tamago` on 8MPLUSLPD4-EVK
// boards.
//
// This package is only meant to be used with `GOOS=tamago` as supported by the
// TamaGo framework for bare metal Go, see https://github.com/usbarmory/tamago.
package goos

import (
	"unsafe"

	"github.com/usbarmory/tamago/arm64"
	"github.com/usbarmory/tamago/board/nxp/imx8mpevk"
	"github.com/usbarmory/tamago/soc/nxp/imx8mp"
)

var (
	RamStart       uint = imx8mp.RamStart
	RamSize        uint = imx8mpevk.RamSize
	RamStackOffset uint = arm64.RamStackOffset

	Bloc   uintptr
	Exit   func(code int32)
	Idle   func(code int64)
	ProcID func() uint64
	Task   func(sp, mp, gp, fn unsafe.Pointer)

	Hwinit0       = arm64.Init
	InitRNG       = imx8mp.InitRNG
	Nanotime      = imx8mp.Nanotime
	GetRandomData = imx8mp.GetRandomData
	Printk        = imx8mpevk.Printk
	Hwinit1       = imx8mpevk.Init
)
