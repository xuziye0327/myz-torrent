package downloader

import (
	"sync"
	"time"

	"github.com/anacrolix/torrent"
)

type magnetItem struct {
	itemName string
	info     torrent.InfoHash
	s        *state
	cTime    time.Time
	uTime    time.Time

	mut          sync.RWMutex
	p            *magnetDownloader
	runningState runningState
}

func (item *magnetItem) id() string {
	return item.info.String()
}

func (item *magnetItem) name() string {
	return item.itemName
}

func (item *magnetItem) createTime() time.Time {
	return item.cTime
}

func (item *magnetItem) start() {
	item.mut.Lock()
	defer item.mut.Unlock()

	if item.runningState == run || item.runningState == finished {
		return
	}

	item.runningState = run
	t, ok := item.p.cli.AddTorrentInfoHash(item.info)
	go func() {
		if ok {
			<-t.GotInfo()
		}
		t.DownloadAll()
	}()
}

func (item *magnetItem) pause() {
	item.mut.Lock()
	defer item.mut.Unlock()

	if item.runningState == pause || item.runningState == finished {
		return
	}

	item.delete()
	item.runningState = pause
}

func (item *magnetItem) delete() {
	item.mut.Lock()
	defer item.mut.Unlock()

	if t, ok := item.p.cli.Torrent(item.info); ok {
		t.Drop()
	}
}

func (item *magnetItem) state() state {
	item.mut.Lock()
	defer item.mut.Unlock()

	t, _ := item.p.cli.Torrent(item.info)
	if t.Info() == nil {
		return state{}
	}

	item.itemName = t.Name()

	totalBytes := int64(0)
	completeBytes := int64(0)
	fs := t.Files()
	for _, f := range fs {
		ps := f.State()

		totalBytes += f.Length()
		for i := range ps {
			if ps[i].Complete {
				completeBytes += ps[i].Bytes
			}
		}
	}

	now := time.Now()
	rate := rate(completeBytes-item.s.CompleteBytes, now.Sub(item.uTime))
	percent := percent(completeBytes, totalBytes)

	s := state{
		TotalBytes:    totalBytes,
		CompleteBytes: completeBytes,
		Rate:          rate,
		Percent:       percent,
	}
	if totalBytes == completeBytes {
		s.RunningState = finished
		t.Drop()
	}

	item.uTime = now
	return s
}

func rate(delta int64, duration time.Duration) float64 {
	if delta < 0 || duration <= 0 {
		return 0
	}

	return float64(delta) / duration.Seconds()
}

func percent(complete, total int64) float64 {
	if total == 0 {
		return 0
	}

	return float64(complete) / float64(total) * 100
}