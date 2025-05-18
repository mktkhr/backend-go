package config

import (
	"log/slog"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	JWT    JWTConfig    `yaml:"jwt"`
	OAuth  OAuthConfig  `yaml:"oauth"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type JWTConfig struct {
	Secret     string `yaml:"secret"`
	Expiration int    `yaml:"expiration"`
}

type OAuthConfig struct {
	ClientID          string `yaml:"clientId"`
	ClientSecret      string `yaml:"clientSecret"`
	RedirectURL       string `yaml:"redirectUrl"`
	SignInCallbackUrl string `yaml:"signInCallbackUrl"`
	OauthDomain       string `yaml:"oauthDomain"`
}

var (
	configInstance *Config
	once           sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		file, err := os.Open("./config/config.yaml")

		if err != nil {
			slog.Error("Configの読み取りに失敗しました。", "error", err.Error())
			os.Exit(1)
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(configInstance); err != nil {
			slog.Error("Configのデコードに失敗しました。", "error", err.Error())
			os.Exit(1)
		}

		slog.Info("Configが正常にロードされました。")
	})

	return configInstance
}
