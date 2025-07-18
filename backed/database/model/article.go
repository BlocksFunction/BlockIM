package model

import (
	"database/sql"
	"gorm.io/datatypes"
)

type Article struct {
	ID       int64          `gorm:"primaryKey;autoIncrement;column:id"`
	Title    string         `gorm:"uniqueIndex;size:255;not null;column:title"`
	Excerpt  string         `gorm:"size:255;not null;column:excerpt"`
	Author   sql.NullString `gorm:"size:255;not null;column:author"`
	Tags     datatypes.JSON `gorm:"type:json;column:tags"`
	ReadTime int64          `gorm:"default:0;column:read_time"`
	Likes    int64          `gorm:"default:0;column:likes"`
	Views    int64          `gorm:"default:0;column:views"`
	Category string         `gorm:"size:255;not null;column:category"`
	Featured bool           `gorm:"default:0;column:featured"`
}
