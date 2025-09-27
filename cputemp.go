package cputemp

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yusufpapurcu/wmi"
)

var log = logrus.New()

type wmiThermalSensor struct {
	InstanceName       string
	CurrentTemperature uint32
}

// GetCPUTemperature returns the CPU temperature from the first ACPI thermal sensor
func GetCPUTemperature() (float64, error) {
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
		return 0.0, fmt.Errorf("unable to get CPU temperature via WMI ACPI: %v", err)
	}

	if len(sensors) == 0 {
		log.Error("No thermal sensors found")
		return 0.0, fmt.Errorf("unable to get CPU temperature via WMI ACPI: no sensors found")
	}

	// Use the first sensor (usually CPU)
	sensor := sensors[0]

	// Convert from Kelvin*10 to Celsius
	kelvin := float64(sensor.CurrentTemperature) / 10.0
	celsius := kelvin - 273.15

	log.Infof("CPU temperature: %.1f°C (sensor: %s)", celsius, sensor.InstanceName)

	return celsius, nil
}
