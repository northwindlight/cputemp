//go:build linux

package cputemp

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func getCPUTemperature() (float64, error) {
	data, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		logrus.Error("Open /sys/class/thermal/thermal_zone0/temp fault:", err)
		return 0.0, fmt.Errorf("Open /sys/class/thermal/thermal_zone0/temp fault: %v", err)
	} else {
		str := strings.TrimSpace(string(data))
		rawdata, err := strconv.Atoi(str)
		if err != nil {
			return 0, fmt.Errorf("解析Linux CPU频率失败: %w", err)
		}
		celsius := float64(rawdata / 1000.0)
		return celsius, nil
	}

}
