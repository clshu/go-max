package iomanager

// IOMangager interface defines the methods for reading and writing data
type IOMangager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
