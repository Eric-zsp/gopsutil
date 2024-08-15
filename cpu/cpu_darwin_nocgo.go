// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && !cgo

package cpu

import "gits.joyconn.cn/go-eric/gopsutil/v4/internal/common"

func perCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}

func allCPUTimes() ([]TimesStat, error) {
	return []TimesStat{}, common.ErrNotImplementedError
}
