package common

type Payload interface {
	Marshal() ([]byte, error)
	Display()
}
