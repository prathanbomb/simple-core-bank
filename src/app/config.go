package app

type Config struct{}

func InitConfig() (*Config, error) {
	config := &Config{}

	return config, nil
}
