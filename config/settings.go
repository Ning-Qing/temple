package config

import (
	"encoding/json"
)

var (
	GlobalSettings *Settings
)

type Settings struct {
	Server ServerSetting `json:"server"`
	Log    LogSetting    `json:"log"`
}

func (s *Settings) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

type ServerSetting struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type LogSetting struct {
	Level string `json:"level"`
}
