package config

import (
	"errors"
	"os"
)

type EnvConfig struct {
	TelegramToken string
	AWSRegion     string
	SQSURL        string
}

func LoadFromEnv() (*EnvConfig, error) {
	cfg := &EnvConfig{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		AWSRegion:     os.Getenv("AWS_REGION"),
		SQSURL:        os.Getenv("SQS_URL"),
	}

	if err := cfg.ValidateWithDefaults(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (e *EnvConfig) ValidateWithDefaults() error {
	if e.TelegramToken == "" {
		return errors.New("TELEGRAM_TOKEN não pode ser vazio")
	}
	if e.AWSRegion == "" {
		return errors.New("AWS_REGION não pode ser vazio")
	}
	if e.SQSURL == "" {
		return errors.New("SQS_URL não pode ser vazio")
	}
	return nil
}
