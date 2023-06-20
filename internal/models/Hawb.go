package models

import (
	"time"
)

type Hawb struct {
	ID          uint
	Origin      string
	Destination string
	Consignor   string
	Consignee   string
	Content     string
	Pieces      int
	Number      string
	Weight      int
	MawbID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
