package config

import (
	"bytes"
	"errors"

	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Version      string
	OneHandedMap Map
	TwoHandedMap Map
	General      GeneralOptions
	//Tapmap       map[uint8]int
}

type GeneralOptions struct {
	SupportMultipleTaps bool
	ComboDelay          int
	TapDevice1Address   string
	TapDevice2Address   string
}

type Map struct {
	Layers []Layer
}

type Layer struct {
	Name            string
	Default         bool
	Trasparent      bool
	SingleTapMap    []TapBinding
	DoubleTapMap    []TapBinding
	TripleTapMap    []TapBinding
	QuadrupleTapMap []TapBinding
}

type TapBinding struct {
	TapCode string
	Action  TapAction
}

type TapActionType string

const (
	KeyPress         TapActionType = "KeyPress"
	MultiKeyPress    TapActionType = "MultiKeyPress"
	KeyHold          TapActionType = "KeyHold"
	MultiKeyHold     TapActionType = "MultiKeyHold"
	Modifier         TapActionType = "Modifier"
	LayerSwitch      TapActionType = "LayerSwitch"
	TempLayerSwitch  TapActionType = "TempLayerSwitch"
	TimedLayerSwitch TapActionType = "TimedLayerSwitch"
	Macro            TapActionType = "Macro"
	Command          TapActionType = "Command"
	ToggleStandby    TapActionType = "ToggleStandby"
)

type TapAction struct {
	ActionType string `toml:"type"`
	Value      string
	Duration   int
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		return nil, errors.New("Config file does not exist")
	} else if err != nil {
		return nil, err
	}

	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func SaveConfig(config Config) {
	var modifyConf bytes.Buffer
	err := toml.NewEncoder(&modifyConf).Encode(config)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("confout.toml", modifyConf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func (config *Config) ProcessConfig() {
	//for i := range config.OneHandedMap.Bindings {
	//_ := config.OneHandedMap.Bindings[i].TapCode

	//}
}
