package hive

import (
	"time"
)

type User struct {
	ID1       uint64
	Name      string
	Age       int64
	Active    bool
	Salary    float32
	CreatedAt time.Time
	UpdatedAt time.Time
	Date      time.Time      `gorm:"type:Date"`
	Score     map[string]int `gorm:"type:MAP<STRING,Int>;serializer:map"`
}
