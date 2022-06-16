package conv

import (
	"fmt"
	"os"
	"path/filepath"
	pluginLoader "plugin"
	"strings"
)

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
		pluginImpl, err := loadPlugin(pluginDir, pluginName)
		if err != nil {
			fmt.Printf("Failed to load plugin %s; Reason: %s\n", pluginName, err)
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

func (pluginStorage *PluginStorage) SendToPlugin(input *ConvInput) (*ConvOutput, error) {
	plugin, err := pluginStorage.findPluginForUnit(input.FromUnit)
	if err != nil {
		return nil, err
	}

	pluginFlags := make(map[string]string)
	pluginIn := input.ToPluginInput(pluginFlags)
	pluginOut, err := plugin.Run(pluginIn)
	if err != nil {
		return nil, err
	}

	return &ConvOutput{
		Value:   pluginOut.Value,
		Unit:    pluginOut.Unit,
		Context: plugin.Context(),
	}, nil
}

func (pluginStorage *PluginStorage) findPluginForUnit(unit string) (ConvPlugin, error) {
	context, exists := pluginStorage.Units[lower(unit)]
	if !exists {
		return nil, fmt.Errorf("cannot find a plugin for %s unit", unit)
	}
	return pluginStorage.Plugins[context], nil
}

func loadPlugin(pluginDir, pluginName string) (ConvPlugin, error) {
	plugin, err := pluginLoader.Open(pluginName)
	if err != nil {
		return nil, err
	}

	pluginImpl, err := plugin.Lookup("PluginImpl")
	if err != nil {
		return nil, err
	}

	var convPlugin ConvPlugin
	convPlugin, ok := pluginImpl.(ConvPlugin)
	if !ok {
		return nil, fmt.Errorf("unexpected type from module symbol")
	}

	return convPlugin, nil
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
