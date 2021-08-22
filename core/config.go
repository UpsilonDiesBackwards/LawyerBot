package core

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfigFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	AppContext.Config = &Config{}
	err = json.Unmarshal(data, AppContext.Config)
	if err != nil {
		return err
	}
	return nil
}
