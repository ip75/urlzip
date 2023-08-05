package core

type IUrlZip interface {
	Validate(string) bool
	ComposeShortURL(string) (string, error)
	GetOriginalURL(string) (string, error)
}
