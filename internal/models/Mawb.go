package models

import (
	"time"
)

type Mawb struct {
	ID          uint
	Number      string
	Origin      string
	Destination string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
