package conv

import "fmt"

type Output struct {
	Value   string
	Unit    string
	Context string
}

func (out Output) String() string {
	return fmt.Sprintf("%s (%s)", out.Value, out.Unit)
}
