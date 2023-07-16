package common

import (
	"github.com/rs/zerolog"
)

type Payload interface {
	Marshal() ([]byte, error)
	Display(*zerolog.Logger)
}
