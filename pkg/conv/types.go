package conv

import "fmt"

type Input struct {
	Value    string
	FromUnit string
	ToUnit   string
	Context  string
}

type Output struct {
	Value   string
	Unit    string
	Context string
}

type ConvPlugin interface {
	Run(in *Input) (*Output, error)
	Context() string
	Units() []string
}

func (in Input) String() string {
	return fmt.Sprintf("%s %s to %s [%s]", in.Value, in.FromUnit, in.ToUnit, in.Context)
}

func (out Output) String() string {
	return fmt.Sprintf("%s %s [%s]", out.Value, out.Unit, out.Context)
}
