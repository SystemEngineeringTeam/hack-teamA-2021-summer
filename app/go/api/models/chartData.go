package models

import "github.com/jinzhu/gorm"

// フィールド名は大文字でなければ外部パッケージからアクセスできないため大文字にすること！
type ChartData struct {
	gorm.Model // データベースで使えるようにするために必要
	UserUID    string
	Spin       int
	Coin       int
}
