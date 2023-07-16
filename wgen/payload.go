package wgen

type Payload interface {
	Marshal() ([]byte, error)
	Display()
}
