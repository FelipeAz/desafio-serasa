package interfaces

// Router eh a interface do ruter.
type Router interface {
	Dispatch(SQLHandler)
}
