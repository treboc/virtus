package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Path       string `yaml:"path"`
		BackupPath string `yaml:"backup_path"`
	} `yaml:"database"`
	Display struct {
		DateFormat string `yaml:"date_format"`
		TimeFormat string `yaml:"time_format"`
	} `yaml:"display"`
}

func DefaultConfig() *Config {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	cfg := &Config{}

	// Database defaults
	cfg.Database.Path = filepath.Join(home, ".virtus", "virtus.db")
	cfg.Database.BackupPath = filepath.Join(home, ".virtus", "backup")

	// Display defaults
	cfg.Display.DateFormat = "2006-01-02"
	cfg.Display.TimeFormat = "15:04:05"

	return cfg
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (c *Config) EnsurePaths() error {
	dbDir := filepath.Dir(c.Database.Path)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return err
	}

	backupDir := filepath.Dir(c.Database.BackupPath)
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return err
	}

	return nil
}
