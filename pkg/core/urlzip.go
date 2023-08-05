package core

import (
	"github.com/ip75/urlzip/pkg/repository"
	"github.com/spaolacci/murmur3"
	"github.com/sqids/sqids-go"
)

const magicNumber = 108

type UrlZipCore struct {
	repository repository.IUZRepository
	sq         *sqids.Sqids
}

func NewUrlZipCore(repo repository.IUZRepository) *UrlZipCore {
	availableCharactersInURL := "0123456789abcdefghijklmnopqrstuvwxyz"
	sqid, _ := sqids.NewCustom(sqids.Options{Alphabet: &availableCharactersInURL})
	return &UrlZipCore{
		repository: repo,
		sq:         sqid,
	}
}

func (c UrlZipCore) Validate(shortPath string) bool {
	numbers := c.sq.Decode(shortPath)

	// check if it valid url without accessing database
	if len(numbers) == 0 || numbers[0] != magicNumber {
		return false
	}
	return true
}

func (c UrlZipCore) ComposeShortURL(fullPath string) (string, error) {

	h := murmur3.New32WithSeed(magicNumber)
	_, _ = h.Write([]byte(fullPath))
	shortPath, err := c.sq.Encode([]uint64{magicNumber, uint64(h.Sum32())})
	if err != nil {
		return "", err
	}

	err = c.repository.Save(fullPath, shortPath)
	if err != nil {
		return "", err
	}
	return shortPath, nil
}

func (c UrlZipCore) GetOriginalURL(shortPath string) (string, error) {
	return c.repository.GetFullURL(shortPath)
}
