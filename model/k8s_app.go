package model

import (
	"gorm.io/gorm"
)

type App struct {
	gorm.Model
	Name      string     `gorm:"index;size(128)" json:"name"`
	Namespace *Namespace `gorm:"column(NamespaceID);rel(fk);comment:'名空间ID'" json:"namespace"`

	MetaData    string `gorm:"type(text)" json:"metaData,omitempty"`
	Description string `gorm:"null;size(512)" json:"description,omitempty"`

	User     string `gorm:"size(128)" json:"user,omitempty"`
	Migrated bool   `gorm:"default(false)" json:"migrated,omitempty"`
}
