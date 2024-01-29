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

5. dbコンテナに入り方
    - `$ docker compose exec -it db bash`

6. appコンテナに入り方＆アプリを起動方法
    - `$ docker compose exec -it app sh`
    - `$ go run main.go`
