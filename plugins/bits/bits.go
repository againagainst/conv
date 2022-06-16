package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/againagainst/conv/pkg/conv"
)

type BitsPlugin struct{}

func (p BitsPlugin) Run(in *conv.Input) (*conv.Output, error) {
	toUnit := "mb"
	if in.ToUnit != "default" {
		toUnit = in.ToUnit
	}

	value, err := BytesConverter(in.Value, in.FromUnit, toUnit)
	if err != nil {
		return nil, err
	}

	output := &conv.Output{
		Value:   value,
		Unit:    toUnit,
		Context: "bits",
	}
	return output, err
}

func (p BitsPlugin) Context() string {
	return "bits"
}

func (p BitsPlugin) Units() []string {
	// ADD:
	// "pb", "kib", "mib"
	return []string{
		"byte",
		"kb",
		"mb",
		"gb",
		"tb",
		"bit",
		"kbit",
		"mbit",
		"gbit",
	}
}

var PluginImpl BitsPlugin

func BytesConverter(value, fromUnit, toUnit string) (string, error) {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value, err
	}

	bits, err := toBits(val, strings.ToLower(fromUnit))
	if err != nil {
		return value, err
	}

	resultFloat64, err := fromBits(bits, strings.ToLower(toUnit))
	if err != nil {
		return value, err
	}

	resultStr := strconv.FormatFloat(resultFloat64, 'g', 15, 64)
	return resultStr, nil
}

var toBitsMult = map[string]float64{
	"byte": 8,
	"kb":   8 * 1024,
	"mb":   8 * 1024 * 1024,
	"gb":   8 * 1024 * 1024 * 1024,
	"tb":   8 * 1024 * 1024 * 1024 * 1024,
	"bit":  1,
	"kbit": 1024,
	"mbit": 1024 * 1024,
	"gbit": 1024 * 1024 * 1024,
}

func toBits(val float64, fromUnit string) (float64, error) {
	mult, ok := toBitsMult[fromUnit]
	if !ok {
		return 0, fmt.Errorf("unit %s is not supported", fromUnit)
	}
	return val * mult, nil
}

func fromBits(bits float64, toUnit string) (float64, error) {
	mult, ok := toBitsMult[toUnit]
	if !ok {
		return 0, fmt.Errorf("unit %s is not supported", toUnit)
	}
	return bits / mult, nil
}
