# go-dev-template

GO + PostgreSQL での開発を、Docker と Air により快適なものとするためのテンプレートリポジトリです。

Go で開発する API サーバー用のコンテナと、PostgreSQL による DB 用のコンテナを Docker Compose により立ち上げます。また、Air によるホットリロードのおかげで、Go のソースコードに対する変更は、立ち上げている API サーバー用のコンテナに反映されます。

詳しい説明については、[こちらの記事](https://qiita.com/KakinokiKanta/items/b943b8aefc13d055aa85)を参照してください。

## 環境構築

本リポジトリを用いて Go を試してみる場合には、以下の手順通りに環境構築してください。
もし、テンプレートとしてリポジトリを作成する場合には、go.mod 内のモジュール名や、docker コンテナ名などは適宜修正してお使いください。

1. [Docker](https://www.docker.com/ja-jp/get-started/)をインストール
2. 本リポジトリをクローンするか、本リポジトリをテンプレートとして選択してリポジトリを作成する
3. `docker network create api-db-network`
4. `docker network create db-pgadmin-network`
5. `docker compose build`

## 開発の流れ

下記コマンドで docker コンテナを起動する。これにより、開発中の API サーバー、Postgres の DB コンテナ、データベースを GUI 操作できる PGAdmin、の 3 つのコンテナが立ち上がる

```
docker compose up
```

docker コンテナを閉じる

```
docker compose down
```

テンプレートでは、API サーバは `localhost:8080` 、 PgAdmin は `localhost:81` で立ち上がる

## 使用技術

テンプレートでは、以下のライブラリやフレームワークを使用していますが、開発者各自の判断で使いたいライブラリ等は追加してください。

- Go
- Gin
- Air
- PostgreSQL
- PgAdmin
- Docker
