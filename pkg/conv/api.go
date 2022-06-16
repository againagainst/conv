package conv

type PluginInput struct {
	Value    string
	FromUnit string
	ToUnit   string
	Flags    map[string]string
}

type PluginOutput struct {
	Value string
	Unit  string
}

type ConvPlugin interface {
	Run(in *PluginInput) (*PluginOutput, error)
	Context() string
	Units() []string
}
