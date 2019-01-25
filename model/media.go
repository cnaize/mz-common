package model

import "github.com/jinzhu/gorm"

type MediaExt string

const (
	MediaExtUnknown MediaExt = ".unknown"
	MediaExtMP3     MediaExt = ".mp3"
)

var SupportedMediaTypes = []MediaExt{MediaExtMP3}

type MediaRoot struct {
	gorm.Model
	Dir        string  `json:"dir" binding:"required" gorm:"not_null;unique_index"`
	ItemsCount uint    `json:"itemsCount"`
	Error      *Error  `json:"error,omitempty"`
	MediaItems []Media `json:"mediaItems,omitempty"`
}

type MediaRootList struct {
	Items []*MediaRoot `json:"items"`
	Error *Error       `json:"error,omitempty"`
}

type Media struct {
	gorm.Model
	Name        string   `json:"name" gorm:"not_null"`
	Ext         MediaExt `json:"ext" gorm:"not_null"`
	Dir         string   `json:"dir" gorm:"not_null"`
	Path        string   `json:"-" gorm:"not_null"` // lower path used for search
	Error       *Error   `json:"error,omitempty"`
	MediaRootID uint     `json:"mediaRootID" gorm:"not_null"`
}

type MediaList struct {
	Items []*Media `json:"items"`
	Error *Error   `json:"error,omitempty"`
}
