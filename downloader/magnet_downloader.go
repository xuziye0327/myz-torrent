package downloader

import (
	"strings"
	"time"

	"github.com/anacrolix/log"
	"github.com/anacrolix/torrent"
)

type magnetDownloader struct {
	cli *torrent.Client
}

func createMangerDownloader(downloadDir string) (downloader, error) {
	c := torrent.NewDefaultClientConfig()
	c.DataDir = downloadDir
	c.NoUpload = true
	c.TorrentPeersHighWater = 200
	c.Logger = log.Discard

	cli, err := torrent.NewClient(c)
	if err != nil {
		return nil, err
	}

	return &magnetDownloader{
		cli: cli,
	}, nil
}

func (downloader *magnetDownloader) new(link string) (downloadItem, error) {
	t, err := downloader.cli.AddMagnet(link)
	if err != nil {
		return nil, err
	}
	// we just want t.InfoHash()
	defer t.Drop()

	return &magnetItem{
		itemName: t.Name(),
		info:     t.InfoHash(),
		cTime:    time.Now(),
		uTime:    time.Now(),

		p:            downloader,
		s:            state{},
		runningState: new,
	}, nil
}

func (downloader *magnetDownloader) validate(link string) bool {
	return strings.HasPrefix(link, "magnet")
}
