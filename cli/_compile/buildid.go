package main

import (
	"os"

	"github.com/virzz/logger"
	"gopkg.in/yaml.v3"
)

const BuildIDFile = ".build_id"

type BuildIDMap map[string]int

func (b BuildIDMap) Get(key string) int {
	v, ok := b[key]
	if ok {
		return v
	}
	return 0
}

func (b BuildIDMap) Inc(key string) (int, error) {
	_, ok := b[key]
	if ok {
		b[key]++
	} else {
		b[key] = 1
	}
	return b[key], b.Save()
}

func (b BuildIDMap) Save() error {
	data, err := yaml.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(BuildIDFile, data, 0775)
}

func (b BuildIDMap) Load() error {
	data, err := os.ReadFile(BuildIDFile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &b)
}

func NewBuildID() BuildIDMap {
	b := BuildIDMap{}
	if err := b.Load(); err != nil {
		logger.ErrorF("Load build id fail: %v", err)
	}
	return b
}

var BuildID = NewBuildID()
