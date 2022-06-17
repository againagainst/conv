package conv

import (
	"fmt"
	"os"
)

func Execute() {
	input := ParseArgs()
	if input == nil {
		PrintUsage()
		os.Exit(0)
	}

	pluginDir := GetPluginDir()
	plugins := LoadPlugins(pluginDir)
	output, err := plugins.SendToPlugin(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	PrintOut(input, output)
	os.Exit(0)
}
