package common

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

// TorrentManager manage all torrent
type TorrentManager struct {
	mut      sync.RWMutex
	cli      *torrent.Client
	torrents map[metainfo.Hash]*Torrent

	dataDir        string
	running        map[metainfo.Hash]*Torrent
	queueing       map[metainfo.Hash]*Torrent
	maxRunningSize int
}

const defaultMaxRunningSize int = 5

// Torrent is my torrent struct
type Torrent struct {
	Name  string       `json:"name"`
	State state        `json:"state"`
	Files torrentFiles `json:"files"`

	t          *torrent.Torrent
	updateTime time.Time
}

type Torrents []*Torrent

func (a Torrents) Len() int {
	return len(a)
}

func (a Torrents) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Torrents) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type torrentFile struct {
	Name  string `json:"name"`
	State state  `json:"state"`
}

type torrentFiles []*torrentFile

type state struct {
	RunningState  runningState `json:"running_state"`
	TotalBytes    int64        `json:"total_bytes"`
	CompleteBytes int64        `json:"complete_bytes"`
	Rate          float64      `json:"rate"`
	Percent       float64      `json:"percent"`
}

type runningState int

const (
	loading runningState = iota
	queueing
	running
	stopping
	finished
)

// InitTorrentManager init torrent manager
func InitTorrentManager(c *Config) (*TorrentManager, error) {
	cc := torrent.NewDefaultClientConfig()
	cc.DataDir = c.DownloadDir
	cc.NoUpload = true
	cc.TorrentPeersHighWater = 200

	cli, err := torrent.NewClient(cc)
	if err != nil {
		return nil, err
	}

	mg := &TorrentManager{
		mut:            sync.RWMutex{},
		cli:            cli,
		torrents:       map[metainfo.Hash]*Torrent{},
		running:        map[metainfo.Hash]*Torrent{},
		queueing:       map[metainfo.Hash]*Torrent{},
		maxRunningSize: defaultMaxRunningSize,
	}

	go func() {
		for {
			mg.updateTorrents()
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			mg.scheduleTorrents()
			time.Sleep(5 * time.Second)
		}
	}()

	return mg, nil
}

// Torrents all torrents
func (mg *TorrentManager) Torrents() Torrents {
	mg.mut.RLock()
	defer mg.mut.RUnlock()

	v := Torrents{}
	for _, t := range mg.torrents {
		v = append(v, t)
	}

	sort.Sort(v)
	return v
}

// AddMagnet add Magnet
func (mg *TorrentManager) AddMagnet(m string) error {
	t, err := mg.cli.AddMagnet(m)
	if err != nil {
		return err
	}
	mg.addTorrent(t)

	// TODO: config auto start
	if err := mg.startTorrent(t.InfoHash()); err != nil {
		return err
	}
	return nil
}

func (mg *TorrentManager) addTorrent(t *torrent.Torrent) {
	mg.mut.Lock()
	defer mg.mut.Unlock()

	go func() {
		<-t.GotInfo()
		fmt.Println("GotInfo: ", t.Name())
	}()

	tt, ok := mg.torrents[t.InfoHash()]
	if !ok {
		tt = &Torrent{
			Name: t.Name(),

			t:          t,
			updateTime: time.Now(),
		}
		mg.torrents[t.InfoHash()] = tt
	}
	tt.updateTorrent()
}

func (mg *TorrentManager) startTorrent(h metainfo.Hash) error {
	mg.mut.Lock()
	defer mg.mut.Unlock()

	t, ok := mg.torrents[h]
	if !ok {
		return fmt.Errorf("Not found torrent: %v ", h)
	}

	if t.State.RunningState == queueing || t.State.RunningState == running || t.State.RunningState == finished {
		return nil
	}

	if mg.maxRunningSize < len(mg.running) {
		mg.running[t.t.InfoHash()] = t
	} else {
		mg.queueing[t.t.InfoHash()] = t
	}

	return nil
}

func (mg *TorrentManager) updateTorrents() {
	mg.mut.Lock()
	defer mg.mut.Unlock()

	for _, t := range mg.torrents {
		t.updateTorrent()
	}
}

func (mg *TorrentManager) scheduleTorrents() {
	mg.mut.Lock()
	defer mg.mut.Unlock()

	for _, t := range mg.torrents {
		if t.State.RunningState == finished {
			delete(mg.running, t.t.InfoHash())
			continue
		}

		if t.State.RunningState != running && t.t.Info() != nil {
			t.State.RunningState = running
			t.t.DownloadAll()
		}
	}

	for len(mg.running) < mg.maxRunningSize && len(mg.queueing) > 0 {
		for _, t := range mg.queueing {
			mg.running[t.t.InfoHash()] = t
			delete(mg.queueing, t.t.InfoHash())
			break
		}
	}

	for _, t := range mg.queueing {
		t.State.RunningState = queueing
	}
}

func (torrent *Torrent) updateTorrent() {
	t := torrent.t
	if t.Info() == nil {
		torrent.State.RunningState = loading
		return
	}
	torrent.Name = t.Name()

	fs := t.Files()
	tfs := torrentFiles{}
	totalBytes := int64(0)
	completeBytes := int64(0)
	for _, f := range fs {
		ps := f.State()

		fTotal := f.Length()
		fComplete := int64(0)
		for i := range ps {
			if ps[i].Complete {
				fComplete += ps[i].Bytes
			}
		}

		totalBytes += fTotal
		completeBytes += fComplete
		tfs = append(tfs, &torrentFile{
			Name: f.Path(),
			State: state{
				TotalBytes:    fTotal,
				CompleteBytes: fComplete,
				Percent:       percent(fComplete, fTotal),
			},
		})
	}
	torrent.Files = tfs

	now := time.Now()
	torrent.State.Rate = rate(completeBytes-torrent.State.CompleteBytes, now.Sub(torrent.updateTime))
	torrent.State.Percent = percent(completeBytes, totalBytes)
	torrent.State.TotalBytes = totalBytes
	torrent.State.CompleteBytes = completeBytes
	if totalBytes == completeBytes {
		torrent.State.RunningState = finished
	}
	torrent.updateTime = now
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
