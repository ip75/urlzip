package repository

type IUZRepository interface {
	Initialize() error
	Save(longPath, shortPath string) error
	GetFullURL(shortPath string) (string, error)
}
