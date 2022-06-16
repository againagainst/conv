package conv

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type ConvInput struct {
	Value    string
	FromUnit string
	ToUnit   string
	Context  string
	Verbose  bool
	Raw      bool
}

func (in *ConvInput) String() string {
	return fmt.Sprintf("%s %s to %s [%s]", in.Value, in.FromUnit, in.ToUnit, in.Context)
}

func (in *ConvInput) ToPluginInput(flags map[string]string) *PluginInput {
	return &PluginInput{
		Value:    in.Value,
		FromUnit: in.FromUnit,
		ToUnit:   in.ToUnit,
		Flags:    flags,
	}
}

func ParseArgs() *ConvInput {
	var context string
	flag.StringVar(&context, "context", "auto", "A plugin name to process unit conversion; default is auto selection")
	var toUnit string
	flag.StringVar(&toUnit, "to", "default", "A result unit; default is determined by plugin")
	verbose := flag.Bool("verbose", false, "show detailed output")
	raw := flag.Bool("raw", false, "show only the value in output")
	if *raw {
		*verbose = false
	}
	flag.Parse()
	arguments := flag.Args()

	return &ConvInput{
		Value:    arguments[0],
		FromUnit: arguments[1],
		ToUnit:   toUnit,
		Context:  context,
		Verbose:  *verbose,
		Raw:      *raw,
	}
}
