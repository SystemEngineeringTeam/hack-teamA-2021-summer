package models

import "github.com/jinzhu/gorm"

// フィールド名は大文字でなければ外部パッケージからアクセスできないため大文字にすること！
type User struct {
	gorm.Model        // データベースで使えるようにするために必要
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Coin       int    `json:"coin"`
}
