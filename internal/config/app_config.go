package config

type AppConfig struct {
	PanelListenAddr string `env:"PANEL_LISTEN_ADDR,notEmpty" envDefault:":3000"`
	MockListenAddr  string `env:"MOCK_LISTEN_ADDR,notEmpty" envDefault:":3001"`
}
