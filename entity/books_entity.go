package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Books struct {
		ID          uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
		Title       string         `json:"title" gorm:"type:varchar(255);not null"`
		Author      string         `json:"author" gorm:"type:varchar(255);not null"`
		Categories  datatypes.JSON `json:"categories" gorm:"type:jsonb;not null"`
		Description string         `json:"description" gorm:"type:text;not null"`
		Publisher   string         `json:"publisher" gorm:"type:varchar(255);not null"`
		PublishDate time.Time      `json:"publish_date" gorm:"type:date;not null"`
		Price       float64        `json:"price" gorm:"type:decimal(10,2);not null"`

		Timestamp
	}
)
