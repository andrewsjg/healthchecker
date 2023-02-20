package healthchecks

type CheckDef struct {
	Description string `mapstructure:"description"`
	Name        string `mapstructure:"name"`
	Check       Check  `mapstructure:"check"`
	Action      Action `mapstructure:"action"`
	Enabled     bool   `mapstructure:"enabled"`
}

type CheckConfig struct {
	Checks []map[string]CheckDef `mapstructure:"checks"`
}

type Check map[string]string

type Action map[string]string

type CheckBlock struct {
	Description string `json:"Description"`
	Name        string `json:"Name"`
	Enabled     bool   `json:"Enabled"`

	Check struct {
		Target string `json:"target"`
		Type   string `json:"type"`
	} `json:"Check"`

	Action struct {
		Type    string `json:"type"`
		Pingurl string `json:"pingurl"`
	} `json:"Action"`
}

/*
type Action struct {
	ActionType string `mapstructure:"type"`
	Pingurl    string `mapstructure:"pingurl"`
} */
