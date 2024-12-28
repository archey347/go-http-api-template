package http

type Config struct {
	Bind    string `mapstructure:"bind"`
	Timeout int    `mapstructure:"timeout"`
}
