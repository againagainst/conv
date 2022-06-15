package conv

import "fmt"

func (output *Output) Sprint() string {
	return fmt.Sprintf("%s (%s)", output.Value, output.Unit)
}
