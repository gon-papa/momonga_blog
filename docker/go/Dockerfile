FROM golang:1.23 as base

WORKDIR /go/src

COPY ../../src/.air.toml ../../src/go.mod ./
# COPY go.mod go.sum ./
RUN go mod download && go mod verify

# TODO: デプロイ時のマルチステージビルドにてCOPYを検討する(Volumeマウントでの開発時は不要)
# COPY ../src/* .
# TODO: デプロイ時のマルチステージビルドにてビルドを検討する
# RUN go build -v -o /usr/local/bin/app ./...

# TODO: 開発のみ使用するライブラリをインストールする
RUN go install github.com/air-verse/air@latest
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

CMD ["air", "-c", ".air.toml"]