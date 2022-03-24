package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not	null;type:varchar(191)"`
	Brand     string `gorm:"not	null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create")

	if len(p.Name) < 4 {
		err = errors.New("Name must be at least 4 characters")
	}
	return
}