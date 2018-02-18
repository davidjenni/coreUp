package machine

// ConfigRenderer is the contract to render a configuration's argument string array
type ConfigRenderer interface {
	Render() ([]string, error)
}