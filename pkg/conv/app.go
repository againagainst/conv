package conv

import (
	"fmt"
	"os"
)

func Execute() {
	input := ParseArgs()
	pluginDir := "./plugins"
	plugins := LoadPlugins(pluginDir)
	output, err := plugins.SendToPlugin(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	PrintOut(input, output)
	os.Exit(1)
}
