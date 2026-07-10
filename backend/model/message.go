package model

import "time"

// Message は掲示板の1件の投稿を表すモデルです。
// この struct の定義をもとに、GORM が messages テーブルを自動で作ります
// (db/db.go の AutoMigrate を参照)。
//
// `json:"..."` タグは、JSONに変換したときのキー名を指定しています。
type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `gorm:"size:255;not null" json:"message"`
	UserName  string    `gorm:"size:255;not null" json:"userName"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
}
