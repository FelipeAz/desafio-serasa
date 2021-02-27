package interfaces

// SQLHandler layer de interface.
type SQLHandler interface {
	CloseConnection() error
}
