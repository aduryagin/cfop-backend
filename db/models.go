package db

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Title       string
	Description string
	Subgroups   []Subgroup `gorm:"foreignkey:GroupID"`
}

type Subgroup struct {
	gorm.Model
	Name         string
	OptimalMoves string
	ImageLink    string
	Algorithms   []Algorithm `gorm:"foreignkey:SubgroupID"`
	GroupID      uint
}

type Algorithm struct {
	gorm.Model
	Algorithm  string
	SubgroupID uint
}
