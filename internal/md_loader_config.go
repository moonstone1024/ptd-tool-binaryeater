package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type MDLoaderConfig struct {
	AddIndex bool `json:"addIndex"`
	Fields   []MDField
}

type MDField struct {
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Include bool      `json:"include"`
	Fields  []MDField `json:"fields"`
}

func LoadMDLoaderConfigFromFile(file *os.File) (*MDLoaderConfig, error) {
	inBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read MD fields file")
	}

	config := &MDLoaderConfig{}
	err = json.Unmarshal(inBytes, config)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode MD fields JSON")
	}

	return config, nil
}
