package testhelper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TestConfig struct {
	Mailings MailingsConfig `json:"mailings"`
}

type MailingsConfig struct {
	TestRecipient string `json:"test_recipient"`
	SendgridConfig SendgridMailingConfig `json:"sendgrid_config"`
}

type SendgridMailingConfig struct {
	ApiKey string `json:"api_key"`
}

func LoadConfig() (*TestConfig, error) {
	gpth := os.Getenv("GOPATH")

	file, err := ioutil.ReadFile(filepath.Join(gpth,"src","github.com","Linus-Boehm","go-serverless-suite", "/config.json"))
	if err != nil {
		return nil, err
	}

	data := TestConfig{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}