package module

import (
	"time"
)

type BaseModule struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
