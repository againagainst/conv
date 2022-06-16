package conv

import "fmt"

type ConvOutput struct {
	Value   string
	Unit    string
	Context string
}

func PrintOut(input *ConvInput, output *ConvOutput) {
	var outputStr string
	if input.Raw {
		outputStr = output.Raw()
	} else if input.Verbose {
		outputStr = output.Verbose()
	} else {
		outputStr = fmt.Sprint(output)
	}
	fmt.Println(outputStr)
}

func (out *ConvOutput) String() string {
	return fmt.Sprintf("%s %s", out.Value, out.Unit)
}

func (out *ConvOutput) Raw() string {
	return fmt.Sprint(out.Value)
}

func (out *ConvOutput) Verbose() string {
	return fmt.Sprintf("%s %s [plugin:%s]", out.Value, out.Unit, out.Context)
}
