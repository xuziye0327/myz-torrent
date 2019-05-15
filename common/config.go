package common

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is a global config using in this proj.
type Config struct {
	DownloadDir string `long:"download-directory" json:"download_directory"`
}

const confName string = ".myz_torrent_config.json"

// LoadConfig will load ~/.myz_torrent_config.json or generate a new config.
func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	confFile := fmt.Sprintf("%v/%v", homeDir, confName)
	f, err := os.Open(confFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				DownloadDir: fmt.Sprintf("%v/myz_torrent_download/", homeDir),
			}, nil
		}
		return nil, err
	}

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	b := make([]byte, info.Size())
	if _, err := f.Read(b); err != nil {
		return nil, err
	}

	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

// SaveConfig will save config to ~/.myz_torrent_config.json.
func (c *Config) SaveConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	confFile := fmt.Sprintf("%v/%v", homeDir, confName)
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Open(confFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	return nil
}
