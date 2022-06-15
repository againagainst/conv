package conv

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

type ConvPlugin interface {
	Run(in *Input) (*Output, error)
	Context() string
	Units() []string
}

type PluginStorage struct {
	Plugins map[string]ConvPlugin
	Units   map[string]string
}

func LoadPlugins(pluginDir string) *PluginStorage {
	pluginStorage := &PluginStorage{
		Plugins: make(map[string]ConvPlugin),
		Units:   make(map[string]string),
	}

	for _, pluginName := range listPlugins(pluginDir) {
		pluginImpl := loadPlugin(pluginDir, pluginName)
		if pluginImpl == nil {
			fmt.Printf("Failed to load plugin %s\n", pluginName)
			continue
		}
		context := lower(pluginImpl.Context())
		pluginStorage.Plugins[context] = pluginImpl
		for _, pluginUnit := range pluginImpl.Units() {
			pluginStorage.Units[lower(pluginUnit)] = context
		}
	}
	return pluginStorage
}

func (pluginStorage *PluginStorage) FindPluginForUnit(unit string) ConvPlugin {
	context, exists := pluginStorage.Units[lower(unit)]
	if !exists {
		return nil
	}
	return pluginStorage.Plugins[context]
}

func loadPlugin(pluginDir, pluginName string) ConvPlugin {
	plug, err := plugin.Open(pluginName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	pluginImpl, err := plug.Lookup("PluginImpl")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var convPlugin ConvPlugin
	convPlugin, ok := pluginImpl.(ConvPlugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		return nil
	}

	return convPlugin
}

func listPlugins(pluginDir string) []string {
	pattern := "*.so"
	var matches []string
	err := filepath.Walk(pluginDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return matches
}

func lower(in string) string {
	return strings.ToLower(in)
}
