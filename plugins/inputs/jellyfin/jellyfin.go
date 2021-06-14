package jellyfin

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

// Jellyfin struct should be named the same as the Plugin
type Jellyfin struct {
	ServerUrl string `toml:"server_url"`
	ApiKey    string `toml:"api_key"`
}

// Description will appear directly above the plugin definition in the config file
func (j *Jellyfin) Description() string {
	return `Retrieves info about a given jellyfin server`
}

// SampleConfig will populate the sample configuration portion of the plugin's configuration
func (j *Jellyfin) SampleConfig() string {
	return `
	## Url of the Jellyfin server
  	server_url = "http://some-ursssl""
  	## Jellyfin API key.
  	# api_key = ""
`
}

// Init can be implemented to do one-time processing stuff like initializing variables
func (j *Jellyfin) Init() error {
	inputs.Add("jellyfin", func() telegraf.Input { return &Jellyfin{} })
	return nil
}

// Gather defines what data the plugin will gather.
func (j *Jellyfin) Gather(acc telegraf.Accumulator) error {
	fields := map[string]interface{}{
		"number": 1,
	}

	acc.AddFields("jellyfin", fields, nil)
	return nil
}
