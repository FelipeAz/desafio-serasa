package interfaces

import "gorm.io/gorm"

// SQLHandler layer de interface.
type SQLHandler interface {
	CloseConnection() error
	GetGorm() *gorm.DB
}
