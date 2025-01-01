# Bulletin Board API - Clean Architecture

このプロジェクトは、Go言語を使用してClean Architectureで設計された掲示板APIです。掲示板の作成・取得が可能なシンプルなAPIを通じて、Clean Architectureの原則を実践しています。

## プロジェクト構成
```
bulletin-board/
├── cmd/
│   └── main.go
├── internal/
│   ├── domain/
│   │   ├── board.go
│   │   └── post.go
│   ├── usecase/
│   │   ├── board_usecase.go
│   │   └── post_usecase.go
│   ├── handler/
│   │   ├── board_handler.go
│   │   └── post_handler.go
│   └── repository/
│       ├── board_repository.go
│       └── post_repository.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── db/
│       └── postgres.go
├── migrations/
│   ├── 001_create_boards_table.up.sql
│   ├── 001_create_boards_table.down.sql
│   ├── 002_create_posts_table.up.sql
│   └── 002_create_posts_table.down.sql
├── scripts/
│   └── migrate.sh
├── .air.toml
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## 各レイヤの役割

### domain
- アプリケーションのビジネスルールを表現します。
- `Board` や `Post` などの構造体を定義します。

### usecase
- ビジネスルールに基づく処理の流れを記述します。
- 例えば、掲示板の作成(`CreateBoard`)や投稿の取得(`GetPost`)などの処理が含まれます。

### handler
- ユーザーからのリクエストを受け付け、`usecase` を呼び出して処理を実行します。
- エンドポイントの定義とリクエストのパラメータ取得、レスポンスの生成を行います。

### repository
- データの永続化を担当します。
- データベース(PostgreSQL)とのやりとりを行い、CRUD操作を実装します。

## 技術スタック
- **言語**: Go
- **Webフレームワーク**: Gorilla mux
- **データベース**: PostgreSQL
- **コンテナ**: Docker, Docker Compose

## 環境構築
### 必要なもの
- Docker
- Docker Compose

### セットアップ
```bash
docker-compose up --build
```

## エンドポイント例
### ボード作成
```
POST /boards
Content-Type: application/json
{
  "name": "General Discussion"
}
```

### ボード一覧取得
```
GET /boards
```

## ライセンス
MIT

