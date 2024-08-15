// SPDX-License-Identifier: BSD-3-Clause
//go:build freebsd

package sensors

import (
	"context"

	"gits.joyconn.cn/go-eric/gopsutil/v4/internal/common"
)

func TemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
