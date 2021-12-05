package main

import (
	"encoding/binary"
	"os"

	"github.com/Elfo404/greenskeeper-observer/internal/logger"
	"gopkg.in/yaml.v3"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var log = logger.NewLogger("main")

func main() {
	config, err := ParseConfig("config.yaml")
	if err != nil {
		log.Panicf("error reading config file",
			"error", err.Error(),
		)
	}

	log.Debug("enabling BLE Stack")

	// Enable BLE interface.
	if err := adapter.Enable(); err != nil {
		log.Panicf("error enabling BLE Stack",
			"error", err.Error(),
		)
	}

	// Start scanning.
	log.Debug("scanning...")
	adapter.Scan(getScanResultHandler(config.Sensors))
}

func readInt16(data []byte) int16 {
	return int16(data[0]&0x7f)<<8 + int16(data[1]) - int16(data[0]&0x80)<<8
}

func getScanResultHandler(sensors Sensors) func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
	sensorCounters := map[string]int{}

	return func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		if result.LocalName() == "BLE Sensor" {
			address := result.Address.String()
			data := result.AdvertisementPayload.GetManufacturerData(0)
			counter := int(data[0])

			temperature := float32(readInt16(data[1:3])) / 100
			humidity := float32(binary.BigEndian.Uint16(data[3:5])) / 100
			soilMoisture := float32(binary.BigEndian.Uint16(data[5:7])) / 100

			alias, exists := sensors[address]
			if !exists {
				alias = address
			}

			if int(counter) != sensorCounters[address] {
				sensorCounters[address] = counter

				log.Infow("received sensor data",
					"address", address,
					"RSSI", result.RSSI,
					"counter", counter,
					"temperature", temperature,
					"humidity", humidity,
					"moisture", soilMoisture,
					"alias", alias,
				)
			}
		}
	}
}

type Sensors map[string]string

type Config struct {
	Sensors Sensors
}

func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
