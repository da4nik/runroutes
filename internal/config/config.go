package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port int `yaml:"port" env:"PORT" env-default:"3333"`
}

func Load(confFilename string) (*Config, error) {
	var conf Config
	err := cleanenv.ReadConfig(confFilename, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
