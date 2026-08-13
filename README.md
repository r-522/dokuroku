# 読録（どくろく）

読録は、読んだものから自分の中に残った感想・考察・解釈をMarkdownとして記録し、あとから探せるようにする個人向けWebアプリケーションです。対象は本だけでなく、技術記事、Web記事、ブログ、複数資料を読んだうえでの考えなどを含みます。

## コンセプト

- 「書く → 残す → あとから探す」に必要な機能だけを持つ。
- タグ、カテゴリ、評価、著者などの入力を必須にしない。
- AIは文章生成ではなく、自由入力を読録のデータ構造へ整理する補助役として使う。
- ユーザー本文は省略・要約・改変せずに保存する。

## 実装状況

- フェーズ1: 設計書作成とCodex設計レビューを完了。
- フェーズ2: Goバックエンド、Reactフロントエンド、PostgreSQLマイグレーション、主要単体テストを実装。

## 技術スタック

| 領域 | 技術 | 方針 |
| --- | --- | --- |
| Frontend | React 19.2.8 + Vite | React標準のstateとブラウザ標準の`fetch`を中心に使う |
| Backend | Go | `net/http`、`database/sql`、`encoding/json`など標準ライブラリ中心 |
| Database | PostgreSQL | Cloud SQL for PostgreSQLを想定 |
| AI | Claude Haiku 4.5 | タイトル判定、Markdown構造化、JSON化の補助に限定 |
| Hosting | GCP | Cloud Run + Cloud SQLを想定 |

## ディレクトリ構成

```text
project-root/
├── README.md
├── 設計書/
├── frontend/
│   ├── index.html
│   ├── package.json
│   └── src/
└── backend/
    ├── cmd/server/
    ├── internal/
    └── migrations/
```

## 起動方法

### 1. PostgreSQLを用意する

`backend/migrations/001_create_book_notes.sql`をPostgreSQLに適用します。

```bash
psql "$DATABASE_URL" -f backend/migrations/001_create_book_notes.sql
```

### 2. Backendを起動する

```bash
cd backend
APP_PASSWORD=1369 \
SESSION_SECRET=local-secret \
DATABASE_URL='postgres://user:password@localhost:5432/dokuroku?sslmode=disable' \
ANTHROPIC_API_KEY='your-key' \
go run ./cmd/server
```

AI APIキーが未設定またはAI呼び出しに失敗した場合でも、本文はフォールバック保存されます。

### 3. Frontendを起動する

```bash
cd frontend
npm install
npm run dev
```

開発時はViteのプロキシ設定を追加するか、同一オリジンで配信する構成に調整します。

## 環境変数

| 変数名 | 用途 |
| --- | --- |
| `DATABASE_URL` | PostgreSQL接続文字列 |
| `APP_PASSWORD` | 簡易パスワード |
| `SESSION_SECRET` | Cookie署名用シークレット。未設定時は`APP_PASSWORD`を使用 |
| `ANTHROPIC_API_KEY` | Claude Haiku 4.5 APIキー |
| `AI_TIMEOUT_SECONDS` | AI処理タイムアウト秒数 |
| `PORT` | Go HTTPサーバーの待受ポート |

`.env`はGitにコミットしません。

## 開発・テスト方法

- Go単体テスト: `cd backend && go test ./...`
- Frontendビルド: `cd frontend && npm install && npm run build`
- Frontendテスト: `cd frontend && npm test`

## 設計書

1. [01_要件定義書](設計書/01_要件定義書.md)
2. [02_基本設計書](設計書/02_基本設計書.md)
3. [03_詳細設計書](設計書/03_詳細設計書.md)
4. [04_単体テスト仕様書](設計書/04_単体テスト仕様書.md)
5. [05_結合テスト仕様書](設計書/05_結合テスト仕様書.md)
6. [06_総合テスト仕様書](設計書/06_総合テスト仕様書.md)
