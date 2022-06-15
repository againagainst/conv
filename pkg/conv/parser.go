package conv

import (
	flag "github.com/spf13/pflag"
)

func ParseArgs() *Input {
	var context string
	flag.StringVar(&context, "context", "auto", "A plugin name to process unit conversion; default is auto selection")
	var toUnit string
	flag.StringVar(&toUnit, "to", "default", "A result unit; default is determined by plugin")
	flag.Parse()
	arguments := flag.Args()

	return &Input{
		Value:    arguments[0],
		FromUnit: arguments[1],
		ToUnit:   toUnit,
		Context:  context,
	}
}
