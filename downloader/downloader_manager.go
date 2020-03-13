package downloader

import (
	"fmt"
	"myz-torrent/common"
	"sort"
	"sync"
	"time"
)

// DownloadManager is DownloadManager
type DownloadManager struct {
	sync.RWMutex

	downloaders        []downloader
	downloadItems      map[string]downloadItem
	downloadItemStates DownloadItemStates
}

// DownloadItemState download item state
type DownloadItemState struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State state  `json:"state"`

	createTime time.Time
}

// DownloadItemStates download item states
type DownloadItemStates []DownloadItemState

// Create a DownloadManager
func Create(c *common.Config) (*DownloadManager, error) {
	mg := &DownloadManager{
		downloadItems: map[string]downloadItem{},
	}

	magnet, err := createMangerDownloader(c.DownloadConfig.DownloadDir)
	if err != nil {
		return nil, err
	}
	mg.downloaders = append(mg.downloaders, magnet)

	go func() {
		for {
			mg.updateState()

			time.Sleep(time.Second)
		}
	}()

	return mg, nil
}

// New download item
func (mg *DownloadManager) New(link string) error {
	mg.Lock()
	defer mg.updateState()
	defer mg.Unlock()

	for _, d := range mg.downloaders {
		if d.validate(link) {
			item, err := d.new(link)
			if err != nil {
				return err
			}

			mg.downloadItems[item.id()] = item

			// TODO: config auto start
			item.start()

			return nil
		}
	}

	return fmt.Errorf("unknow download link")
}

// Strart a download item
func (mg *DownloadManager) Strart(id string) {
	mg.Lock()
	defer mg.updateState()
	defer mg.Unlock()

	item, ok := mg.downloadItems[id]
	if ok {
		item.start()
	}
}

// Pause a download item
func (mg *DownloadManager) Pause(id string) {
	mg.Lock()
	defer mg.Unlock()

	item, ok := mg.downloadItems[id]
	if !ok {
		return
	}

	item.pause()
}

// Delete a download item
func (mg *DownloadManager) Delete(id string) {
	mg.Lock()
	defer mg.updateState()
	defer mg.Unlock()

	item, ok := mg.downloadItems[id]
	if !ok {
		return
	}

	item.delete()
	delete(mg.downloadItems, id)
}

// State returns all states of download item
func (mg *DownloadManager) State() DownloadItemStates {
	mg.RLock()
	defer mg.RUnlock()

	return mg.downloadItemStates
}

func (mg *DownloadManager) updateState() {
	mg.Lock()
	defer mg.Unlock()

	infos := make(DownloadItemStates, 0)
	for _, item := range mg.downloadItems {
		infos = append(infos, DownloadItemState{
			ID:    item.id(),
			Name:  item.name(),
			State: item.state(),

			createTime: item.createTime(),
		})
	}

	sort.Sort(infos)
	mg.downloadItemStates = infos
}

func (a DownloadItemStates) Len() int           { return len(a) }
func (a DownloadItemStates) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DownloadItemStates) Less(i, j int) bool { return a[i].createTime.Before(a[j].createTime) }
