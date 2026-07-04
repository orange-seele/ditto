package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveLayout   string `json:"active_layout"`
	ActiveSize     int    `json:"active_size"`
	ActiveStandard string `json:"active_standard"`
	Locked         bool   `json:"locked"`
	ShowAllInfo    *bool  `json:"show_all_info,omitempty"`
}

const DirName = "ditto"

func configPath() (string, error) {
	cfgDir := os.Getenv("XDG_CONFIG_HOME")
	if cfgDir == "" {
		var err error
		cfgDir, err = os.UserConfigDir()
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(cfgDir, DirName, "config.json"), nil
}

func Default() Config {
	return Config{ActiveLayout: "qwerty", ActiveSize: 75, ActiveStandard: "ansi"}
}

func LoadConfig() Config {
	path, err := configPath()
	if err != nil {
		return Default()
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Default()
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Default()
	}
	if cfg.ActiveLayout == "" {
		cfg.ActiveLayout = Default().ActiveLayout
	}
	if cfg.ActiveSize == 0 {
		cfg.ActiveSize = Default().ActiveSize
	}
	if cfg.ActiveStandard == "" {
		cfg.ActiveStandard = Default().ActiveStandard
	}
	return cfg
}

func SaveConfig(cfg Config) error {
	path, err := configPath()
	if err != nil {
		return fmt.Errorf("config path: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir config dir: %w", err)
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	return nil
}
