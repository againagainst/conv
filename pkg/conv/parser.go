package conv

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type Input struct {
	Value    string
	FromUnit string
	ToUnit   string
	Context  string
}

func (in Input) String() string {
	return fmt.Sprintf("%s %s to %s [%s]", in.Value, in.FromUnit, in.ToUnit, in.Context)
}

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
