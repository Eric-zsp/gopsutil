// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && !cgo

package sensors

import (
	"context"

	"github.com/Eric-zsp/gopsutil/v4/internal/common"
)

func TemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
