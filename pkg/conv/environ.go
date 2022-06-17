package conv

import (
	"os"
)

func GetPluginDir() string {
	dir := os.Getenv("CONV_PLUGINS")
	if dir == "" {
		dir = "/usr/local/lib/conv/plugin"
	}
	return dir
}
