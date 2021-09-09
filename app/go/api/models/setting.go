package models

import "github.com/jinzhu/gorm"

// フィールド名は大文字でなければ外部パッケージからアクセスできないため大文字にすること！
type Setting struct {
	gorm.Model        // データベースで使えるようにするために必要
	Path       string `json:"path" gorm:"type:varchar(1000)"`
	UserUID    string `json:"user_uid"`
}
