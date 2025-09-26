package cputemp

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yusufpapurcu/wmi"
)

var log = logrus.New()

type TemperatureReading struct {
	TemperatureC float64
	TemperatureF float64
	InstanceName string
	Success      bool
}

type wmiThermalSensor struct {
	InstanceName       string
	CurrentTemperature uint32
}

// GetCPUTemperature returns the CPU temperature from the first ACPI thermal sensor
func GetCPUTemperature() (*TemperatureReading, error) {
	// Set up logger
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	var sensors []wmiThermalSensor

	// Query ACPI thermal zone temperature
	err := wmi.QueryNamespace("SELECT InstanceName, CurrentTemperature FROM MSAcpi_ThermalZoneTemperature",
		&sensors, "root\\WMI")

	if err != nil {
		log.Errorf("WMI query failed: %v", err)
		return nil, fmt.Errorf("unable to get CPU temperature via WMI ACPI: %v", err)
	}

	if len(sensors) == 0 {
		log.Error("No thermal sensors found")
		return nil, fmt.Errorf("unable to get CPU temperature via WMI ACPI: no sensors found")
	}

	// Use the first sensor (usually CPU)
	sensor := sensors[0]

	// Convert from Kelvin*10 to Celsius
	kelvin := float64(sensor.CurrentTemperature) / 10.0
	celsius := kelvin - 273.15
	fahrenheit := celsius*1.8 + 32

	log.Infof("CPU temperature: %.1f°C (sensor: %s)", celsius, sensor.InstanceName)

	return &TemperatureReading{
		TemperatureC: celsius,
		TemperatureF: fahrenheit,
		InstanceName: sensor.InstanceName,
		Success:      true,
	}, nil
}

// CheckAdminPermissions verifies if the application has administrator privileges
func CheckAdminPermissions() bool {
	var sensors []wmiThermalSensor

	// Try to query a WMI class that requires admin privileges
	err := wmi.QueryNamespace("SELECT InstanceName FROM MSAcpi_ThermalZoneTemperature",
		&sensors, "root\\WMI")

	if err != nil {
		log.Warnf("Admin check failed: %v - application may need administrator privileges", err)
		return false
	}

	log.Info("Admin permissions verified")
	return true
}
