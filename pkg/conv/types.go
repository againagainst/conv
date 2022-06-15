package conv

import "fmt"

type Input struct {
	value    string
	fromUnit string
	toUnit   string
	context  string
}

type Output struct {
	value   string
	unit    string
	context string
}

func (in Input) String() string {
	return fmt.Sprintf("%s %s to %s [%s]", in.value, in.fromUnit, in.toUnit, in.context)
}

func (out Output) String() string {
	return fmt.Sprintf("%s %s [%s]", out.value, out.unit, out.context)
}
