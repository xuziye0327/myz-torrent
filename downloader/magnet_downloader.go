package downloader

import (
	"strings"
	"sync"
	"time"

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

		mut:          sync.RWMutex{},
		p:            downloader,
		runningState: new,
	}, nil
}

func (downloader *magnetDownloader) validate(link string) bool {
	return strings.HasPrefix(link, "magnet")
}
