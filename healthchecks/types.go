package healthchecks

type CheckDef struct {
	Description string  `mapstructure:"description"`
	Name        string  `mapstructure:"name"`
	Checks      Checks  `mapstructure:"checks"`
	Actions     Actions `mapstructure:"actions"`
	Enabled     bool    `mapstructure:"enabled"`
}

type Healthchecks struct {
	Healthchecks []map[string]CheckDef `mapstructure:"healthchecks"`
}

type Checks map[string]map[string]string

type Actions map[string]map[string]string

// type to map data sent to the API from the frontend
type CheckBlock struct {
	Description string `json:"Description"`
	Name        string `json:"Name"`
	Enabled     bool   `json:"Enabled"`

	Checks Checks `json:"Checks"`

	Actions Actions `json:"Actions"`
}

/*
type Action struct {
	ActionType string `mapstructure:"type"`
	Pingurl    string `mapstructure:"pingurl"`
} */
