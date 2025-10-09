//go:build windows

package cputemp

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yusufpapurcu/wmi"
)

type wmiThermalSensor struct {
	InstanceName       string
	CurrentTemperature uint32
}

func getCPUTemperature() (float64, error) {
	// Set up logger
	var sensors []wmiThermalSensor

	// Query ACPI thermal zone temperature
	err := wmi.QueryNamespace("SELECT InstanceName, CurrentTemperature FROM MSAcpi_ThermalZoneTemperature",
		&sensors, "root\\WMI")

	if err != nil {
		logrus.Errorf("WMI query failed: %v", err)
		return 0.0, fmt.Errorf("unable to get CPU temperature via WMI ACPI: %v", err)
	}

	if len(sensors) == 0 {
		logrus.Error("No thermal sensors found")
		return 0.0, fmt.Errorf("unable to get CPU temperature via WMI ACPI: no sensors found")
	}

	// Use the first sensor (usually CPU)
	sensor := sensors[0]

	// Convert from Kelvin*10 to Celsius
	kelvin := float64(sensor.CurrentTemperature) / 10.0
	celsius := kelvin - 273.15

	//logrus.Infof("CPU temperature: %.1f°C (sensor: %s)", celsius, sensor.InstanceName)

	return celsius, nil
}
