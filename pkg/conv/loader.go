package conv

import (
	"fmt"
	"os"
	"plugin"
)

func LoadPlugin(input *Input) (*Output, error) {
	mod := "./plugins/ip.so"

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pluginImpl, err := plug.Lookup("PluginImpl")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var convPlugin ConvPlugin
	convPlugin, ok := pluginImpl.(ConvPlugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return convPlugin.Run(input)
}
