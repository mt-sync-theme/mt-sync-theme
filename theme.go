package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Theme struct {
	Id        string
	Directory string
}

type ThemeYAML struct {
	Id string
}

func NewTheme(directory string) (Theme, error) {
	content, err := ioutil.ReadFile(path.Join(directory, "theme.yaml"))
	if err != nil {
		return Theme{}, err
	}

	data := ThemeYAML{}
	err = yaml.Unmarshal([]byte(content), &data)
	if err != nil {
		return Theme{}, err
	}

	if data.Id == "" {
		return Theme{}, errors.New("Can not parse theme.yaml")
	}

	theme := Theme{
		Id:        data.Id,
		Directory: directory,
	}

	return theme, nil
}
