package config

import (
	"encoding/json"
	"sync"
)

type Settings struct {
	lock   *sync.RWMutex
	Server ServerSetting `json:"server"`
	Log    LogSetting    `json:"log"`
}

func (s *Settings) GetServerHost() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Server.Host
}

func (s *Settings) GetServerPort() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Server.Port
}

func (s *Settings) GetServerMode() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Server.Mode
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

var (
	DefaultHost = ""
	DefaultPort = 8080
	DefaultMode = "debug"
)
