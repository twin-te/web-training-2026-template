# web-training-2026-template

Twin:te の研修テンプレート。

## 構成

```
backend/   - Go + Echo + GORM (PostgreSQL) + Docker
frontend/  - Vite + React + TypeScript
```

それぞれ独立したプロジェクトです。詳細は各ディレクトリの README を参照。

## クイックスタート

ターミナルを 2 つ用意して、

**backend**

```bash
cd backend
cp .env.example .env
docker compose up --build
```

- API: http://localhost:3000
- Adminer (DB 閲覧): http://localhost:8080 (System: PostgreSQL / Server: db / User: app / Password: app / Database: app)

データベースのテーブルはアプリ起動時に自動で作られます(GORM の AutoMigrate)。

**frontend**

```bash
cd frontend
npm install
npm run dev
```

- http://localhost:5173
