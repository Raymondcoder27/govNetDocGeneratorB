package models

import "gorm.io/gorm"

type PDF struct {
    gorm.Model
    ID   string `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Data string `json:"data"`
}
