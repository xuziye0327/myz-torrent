package common

import "flag"

// Config is a global config using in this proj.
type Config struct {
	ServerAddr   string `string:"server" json:"server"`
	ServerPortal int    `int:"portal" json:"portal"`

	DownloadConfig *DownloadConfig `DownloadConfig:"download_config" json:"download_config"`
}

// DownloadConfig is using in running time
type DownloadConfig struct {
	DownloadDir string `long:"download-directory" json:"download_directory"`
}

// LoadConfig will load ~/.myz_torrent_config.json or generate a new config.
func LoadConfig() (*Config, error) {
	var p int
	flag.IntVar(&p, "p", 8080, "")

	var s string
	flag.StringVar(&s, "s", "0.0.0.0", "")

	var d string
	flag.StringVar(&d, "d", "~/myz_torrent_download/", "")

	var c string
	flag.StringVar(&c, "c", "", "")

	flag.Parse()

	return &Config{
		ServerAddr:   s,
		ServerPortal: p,

		DownloadConfig: &DownloadConfig{
			DownloadDir: d,
		},
	}, nil
}
