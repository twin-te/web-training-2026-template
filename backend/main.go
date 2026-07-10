package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/twin-te/web-training-2026-template/backend/db"
)

func main() {
	e := echo.New() // 1.

	// 2.
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// データベースに接続し、モデルの定義からテーブルを自動生成する
	if err := db.Init(); err != nil {
		e.Logger.Fatal(err)
	}

	/* ここに追記 */

	// 3.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// 4.
	e.Logger.Fatal(e.Start(":" + port))
}
