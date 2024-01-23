## プルリクテンプレート使い方
?template=simple.md

## 環境構築手順
1. リモートレポジトリをクローンする
    - `$ git clone git@github.com:tetsu-777/rakuraku-reserve-app.git`

2. env-exampleファイルをコピーして.envファイルをプロジェクト直下のディレクトリに作成する
    - `$ cp env-example .env`

3. docker image　の作成
    - `$ docker compose build`

4. コンテナ作成＆起動
    - `$ docker compose up -d`

5. dbコンテナに入り、tableを作成
    - `$ docker compose exec -it db bash`
    - `$ psql -U postgres -d app`
    - `$ \i docker-entrypoint-initdb.d/init_database.sql`
    - `$ \q`
    - `$ exit`

6. appコンテナに入り、アプリを起動
    - `$ docker compose exec -it app sh`
    - `$ go run main.go`

7. indexのページを確認
    - [リンク](http://localhost:8080)
