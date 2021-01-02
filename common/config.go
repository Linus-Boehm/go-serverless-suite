package common

import (
	"embed"
	"encoding/json"

	"github.com/pkg/errors"
)

type Configer interface {
	IsDebug() bool
	GetStage() string
	Validate() (valid bool, err error)
}

func LoadConfigFile(filepath string, fs embed.FS, config Configer) (Configer, error) {
	file, err := fs.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	valid, err := config.Validate()
	if err != nil {
		return config, err
	}
	if !valid {
		return config, errors.New("config not validation failed")
	}
	return config, nil
}
