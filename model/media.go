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
	Dir        string  `json:"dir" gorm:"unique_index" binding:"required"`
	ItemsCount uint `json:"itemsCount"`
	Error      *Error  `json:"error,omitempty"`
	MediaItems []Media `json:"mediaItems,omitempty"`
}

type MediaRootList struct {
	Items []MediaRoot `json:"items"`
	Error *Error      `json:"error,omitempty"`
}

type Media struct {
	gorm.Model
	Name        string   `json:"name"`
	Ext         MediaExt `json:"ext"`
	Path string `json:"path"`
	Error       *Error   `json:"error,omitempty"`
	MediaRootID uint     `json:"mediaRootID"`
}

type MediaList struct {
	Items []Media `json:"items"`
	Error *Error  `json:"error,omitempty"`
}
