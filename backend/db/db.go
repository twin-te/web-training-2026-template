package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/twin-te/web-training-2026-template/backend/model"
)

// DB はアプリケーション全体で共有するデータベース接続です。
var DB *gorm.DB

// Init はデータベースに接続し、モデルの定義からテーブルを自動生成(AutoMigrate)します。
func Init() error {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	d, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// model.Message の定義に合わせて messages テーブルを自動で作成・更新する
	if err := d.AutoMigrate(&model.Message{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	DB = d
	return nil
}
