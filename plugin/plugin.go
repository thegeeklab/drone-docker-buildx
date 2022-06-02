package plugin

import (
	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
)

// Plugin implements drone.Plugin to provide the plugin implementation.
type Plugin struct {
	settings Settings
	pipeline drone.Pipeline
	network  drone.Network
}

// New initializes a plugin from the given Settings, Pipeline, and Network.
func New(settings Settings, pipeline drone.Pipeline, network drone.Network) drone.Plugin {
	return &Plugin{
		settings: settings,
		pipeline: pipeline,
		network:  network,
	}
}
