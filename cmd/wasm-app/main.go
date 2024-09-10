package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	extism "github.com/extism/go-sdk"
)

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "./bin/wasm-app-plugin.wasm",
				Name: "main",
			},
		},
		Config:       make(map[string]string),
		AllowedHosts: []string{},
		AllowedPaths: make(map[string]string),
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	exit, out, err := plugin.Call("convert2Sigs", fullV2InputJSON)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}
	signals := []vss.Signal{}
	err = json.Unmarshal(out, &signals)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(signals)

}

var fullV2InputJSON = []byte(`{
    "id": "2fHbFXPWzrVActDb7WqWCfqeiYe",
    "source": "dimo/integration/123",
    "specversion": "1.0",
    "dataschema": "testschema/v2.0",
    "subject": "0x98D78d711C0ec544F6fb5d54fcf6559CF41546a9",
    "time": "2024-04-18T17:20:46.436008782Z",
    "type": "com.dimo.device.status",
    "signature": "0x72208df3282c890ec72afe22abbcffb76ec73dc3e1ce8becd158469126f10c35245289e02ad41782e07376f9b9092a2fec96477a6e453fed1ca3860639e776f31b",
    "data": {
        "timestamp": 1713460846435,
        "device": {
            "rpiUptimeSecs": 218,
            "batteryVoltage": 12.28
        },
        "vehicle": {
            "signals": [
                {
                    "timestamp": 1713460823243,
                    "name": "longFuelTrim",
                    "value": 25
                },
                {
                    "timestamp": 1713460826633,
                    "name": "coolantTemp",
                    "value": 107
                },
                {
                    "timestamp": 1713460827173,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460829314,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460829844,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460830382,
                    "name": "shortFuelTrim",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460837235,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460842256,
                    "name": "maf",
                    "value": 475.79
                },
                {
                    "timestamp": 1713460844422,
                    "name": "engineLoad",
                    "value": 12.54912
                },
                {
                    "timestamp": 1713460844962,
                    "name": "throttlePosition",
                    "value": 23.529600000000002
                },
                {
                    "timestamp": 1713460845497,
                    "name": "shortFuelTrim",
                    "value": 12.5
                },
                {
                    "timestamp": 1713460846435,
                    "name": "isRedacted",
                    "value": false
                },
                {
                    "timestamp": 1713460846435,
                    "name": "longitude",
                    "value": -56.50151833333334
                },
                {
                    "timestamp": 1713460846435,
                    "name": "latitude",
                    "value": 56.27014
                },
                {
                    "timestamp": 1713460846435,
                    "name": "hdop",
                    "value": 1.4
                },
                {
                    "timestamp": 1713460846435,
                    "name": "nsat",
                    "value": 6
                },
                {
                    "timestamp": 1713460846435,
                    "name": "wpa_state",
                    "value": "COMPLETED"
                },
                {
                    "timestamp": 1713460846435,
                    "name": "ssid",
                    "value": "foo"
                },
                {
                    "timestamp": 1713460846435,
                    "name": "vehicleSpeed",
                    "value": 39
                },
                {
                    "timestamp": 1713460846435,
                    "name": "rpm",
                    "value": 2000
                },
                {
                    "timestamp": 1713460846435,
                    "name": "fuelLevel",
                    "value": 50
                },
            ]
        }
    },
    "vehicleTokenId": 123,
    "make": "",
    "model": "",
    "year": 0
}`)
