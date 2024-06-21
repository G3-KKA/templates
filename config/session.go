package config

//

type token string
type configVersion int
type Session struct {
	newConfig                    Config
	tokenThatShouldGetOldConfigs map[token]configVersion
	oldConfigs                   map[configVersion]Config
}
type ConfigHolder interface {
	GetConfig(token) (Config, error)
}
