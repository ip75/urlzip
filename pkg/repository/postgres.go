package repository

import (
	"log"
	"time"

	"gorm.io/gorm/clause"
)

type Relation struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:true"`
	ShortPath string    `gorm:"column:short_path;index;unique;comment:generated short url"`
	FullPath  string    `gorm:"column:full_path;comment:source full url"`
	Created   time.Time `gorm:"<-:create;index;default:current_timestamp"`
}

// Initialize ...
func (d *UrlZipDatabase) Initialize() error {
	return d.AutoMigrate(&Relation{})
}

// Save saves the pair of urls to database
func (d *UrlZipDatabase) Save(fullPath, shortPath string) error {

	result := d.
		Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "short_path"}},
				DoNothing: true,
			}).
		Create(&Relation{ShortPath: shortPath, FullPath: fullPath})

	return result.Error
}

func (d *UrlZipDatabase) GetFullURL(shortPath string) (string, error) {

	var relation Relation
	r := d.Where("short_path = ?", shortPath).Find(&relation)
	if r.RowsAffected == 0 {
		log.Printf("relation is not found %v", r.Error)
	}
	return relation.FullPath, r.Error
}
