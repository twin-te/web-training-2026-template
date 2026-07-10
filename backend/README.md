# backend

Go + Echo + GORM (PostgreSQL) のバックエンドテンプレート。

## スタック

- Go
- [Echo](https://echo.labstack.com/) v4
- [GORM](https://gorm.io/) + PostgreSQL 16
- [Adminer](https://www.adminer.org/) (DB をブラウザから覗ける、docker compose)
- [air](https://github.com/air-verse/air) (ホットリロード、docker compose の dev target)

## 開発の始め方

`.env.example` をコピーして `.env` を作ってから、Docker でまるごと起動します。

```bash
cp .env.example .env
docker compose up --build
```

- API: http://localhost:3000
- Adminer: http://localhost:8080 (System: PostgreSQL / Server: db / User: app / Password: app / Database: app)

データベースのテーブルは、アプリ起動時に GORM の AutoMigrate がモデル定義から自動で作ります。
マイグレーションコマンドを手で打つ必要はありません。

ソースコードを保存すると air が自動でビルドし直してくれます(ホットリロード)。

## ディレクトリ

```
main.go            - エントリポイント (Echo の起動)
api/
  message.go       - /messages エンドポイントのハンドラ
db/
  db.go            - DB 接続 + AutoMigrate
model/
  message.go       - Message モデル (テーブル定義)
Dockerfile
docker-compose.yml
.air.toml          - air (ホットリロード) の設定
```

## サンプル API

```bash
curl http://localhost:3000/health
curl http://localhost:3000/messages
curl -X POST http://localhost:3000/messages \
  -H 'Content-Type: application/json' \
  -d '{"message":"こんにちは!","userName":"風吹けば名無し"}'
```

(研修の Phase を進めるまでは、まだ実装されていないエンドポイントもあります)
