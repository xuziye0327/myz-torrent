package downloader

type downloader interface {
	new(link string) (downloadItem, error)
	validate(link string) bool
}
