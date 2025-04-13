package model

import "gorm.io/gorm"

type FileMetadata struct {
	gorm.Model
	Filename    string `gorm:"size:255;not null;index" json:"filename"`
	Author      string `gorm:"size:255;not null;index" json:"author"`
	ContentType string `gorm:"size:100;not null" json:"content_type"`
	Hash        string `gorm:"size:100;unique;not null;index" json:"hash"`
}
