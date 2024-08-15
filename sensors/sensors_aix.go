// SPDX-License-Identifier: BSD-3-Clause
//go:build aix

package sensors

import (
	"context"

	"gits.joyconn.cn/go-eric/gopsutil/v4/internal/common"
)

const (
	hostTemperatureScale = 1000.0 // Not part of the linked file, but kept just in case it becomes relevant
)

func VirtualizationWithContext(ctx context.Context) (string, string, error) {
	return "", "", common.ErrNotImplementedError
}
