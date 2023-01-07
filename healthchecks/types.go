package healthchecks

type CheckDef struct {
	Description string `mapstructure:"description"`
	Check       Check  `mapstructure:"check"`
	Action      Action `mapstructure:"action"`
}

type CheckConfig struct {
	Checks []map[string]CheckDef `mapstructure:"checks"`
}

type Check map[string]string

type Action map[string]string

/*
type Action struct {
	ActionType string `mapstructure:"type"`
	Pingurl    string `mapstructure:"pingurl"`
} */
