package interfaces

import "gorm.io/gorm"

// SQLHandler layer de interface.
type SQLHandler interface {
	CloseConnection()
	GetGorm() *gorm.DB
}
