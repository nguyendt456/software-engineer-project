FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
