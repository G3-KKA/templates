package config

// Ideas for dynamic config

type token string
type configVersion int
type Session struct {
	newConfig                    config
	tokenThatShouldGetOldConfigs map[token]configVersion
	oldConfigs                   map[configVersion]config
}
type ConfigHolder interface {
	GetConfig(token) (config, error)
}
