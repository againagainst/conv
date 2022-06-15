package conv

import (
	"fmt"
	"os"
)

func Execute() {
	input := ParseArgs()
	pluginDir := "./plugins"
	plugins := LoadPlugins(pluginDir)
	converter := plugins.FindPluginForUnit(input.FromUnit)
	if converter == nil {
		fmt.Printf("Can't find a plugin for unit %s\n", input.FromUnit)
		os.Exit(1)
	}
	result, err := converter.Run(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
	os.Exit(1)
}
