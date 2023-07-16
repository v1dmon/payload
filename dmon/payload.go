package dmon

type Payload interface {
	Marshal() ([]byte, error)
	Display()
}
