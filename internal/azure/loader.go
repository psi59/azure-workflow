package azure

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServicesFile struct {
	Services []Service `yaml:"services"`
}

func LoadServices(path string) ([]Service, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var file ServicesFile
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	return file.Services, nil
}
