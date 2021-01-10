package testhelper

import (
	"embed"
	"encoding/json"
	"io/ioutil"
)

//go:embed assets/*
var configAssets embed.FS

type TestConfig struct {
	Mailings MailingsConfig `json:"mailings"`
}

type MailingsConfig struct {
	TestRecipient  string                `json:"test_recipient"`
	SendgridConfig SendgridMailingConfig `json:"sendgrid_config"`
}

type SendgridMailingConfig struct {
	APIKey string `json:"api_key"`
	ListID string `json:"list_id"`
}

func LoadConfig() (*TestConfig, error) {
	file, err := configAssets.Open("assets/config.json")
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	data := TestConfig{}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
