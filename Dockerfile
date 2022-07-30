#マルチステージビルド
#docker build -t go_todo:任意のタグ --target deploy ./
#デプロイにはリリース用のバイナリファイルのみコピーする
#実用Go言語 14.2を参考にする

# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ---------------------------------------------------
#デプロイ用のコンテナ
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ---------------------------------------------------

#ローカル環境で利用するホットリロード環境
#ファイル変更が検知するたびにbuildコマンドが走り、再起動する
#ローカルのファイルをマウントしておく
#airコマンド用に.air.tomlを用意する
FROM golang:1.18.2 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]