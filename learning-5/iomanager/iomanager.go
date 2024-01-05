package iomanager

type IOManager interface {
	Read() ([]string, error)
	Write(any) error
}
