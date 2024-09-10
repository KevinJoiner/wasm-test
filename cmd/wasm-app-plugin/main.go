//go:build !std
// +build !std

package main

import (
	"context"

	"github.com/DIMO-Network/model-garage/pkg/vss/convert"
	pdk "github.com/extism/go-pdk"
)

func main() {
}

//export convert2Sigs
func convert2Sigs() int32 {
	input := pdk.Input()
	actualSignals, err := convert.SignalsFromPayload(context.Background(), nil, input)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.OutputJSON(actualSignals)

	return 0
}
