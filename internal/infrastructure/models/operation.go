package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lucabecci/concurrent_publisher_go/internal/domain/entities"
	"gorm.io/gorm"
)

type Operation struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Status      entities.Status
	Value       int
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
}
