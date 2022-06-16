package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/againagainst/conv/pkg/conv"
)

type IpPlugin struct{}

func (p IpPlugin) Run(in *conv.PluginInput) (*conv.PluginOutput, error) {
	value, err := Ipv4ToHex(in.Value)
	if err != nil {
		return nil, err
	}
	output := &conv.PluginOutput{
		Value: value,
		Unit:  "hex",
	}
	return output, err
}

func (p IpPlugin) Context() string {
	return "ip"
}

func (p IpPlugin) Units() []string {
	return []string{"hex", "ip"}
}

var PluginImpl IpPlugin

func Ipv4ToHex(ipv4 string) (string, error) {
	ipv4RegexpText := []string{
		"^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\",
		".(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\",
		".(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\",
		".(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$",
	}
	ipv4Regexp := regexp.MustCompile(strings.Join(ipv4RegexpText, ""))
	parsed := ipv4Regexp.FindStringSubmatch(ipv4)

	if parsed == nil {
		return ipv4, errors.New("input does not match ipv4 format")
	}

	result := make([]string, 4)
	for i, strint := range parsed[1:] {
		strhex, err := intToHex(strint)
		if err != nil {
			return strint, err
		}
		result[i] = strhex
	}
	return fmt.Sprintf("0x%s", strings.Join(result, "")), nil
}

func intToHex(strInt string) (strHex string, err error) {
	value, err := strconv.ParseInt(strInt, 10, 64)
	if err == nil {
		strHex = fmt.Sprintf("%02s", strings.ToUpper(strconv.FormatInt(value, 16)))
	}
	return strHex, err
}
