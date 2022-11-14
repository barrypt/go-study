package model

import (
	"time"

	"gorm.io/gorm"
)

type Level int8

const (
	Primary Level = iota
	Middle
	Senior
	Master
)

type User struct {
	Name string
	Level Level
	YearsService  int
	TimeEntry  time.Time
	Score  float32
	Tag string 
	PlatformCertification int8
	AreasExpertise  string
	PersonalProfile  string
	Follower  int32
	ServiceCount int32
	Birthday  time.Time
	Password string
	gorm.Model
}
