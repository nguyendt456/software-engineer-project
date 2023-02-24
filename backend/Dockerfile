FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main server/main.go server/create_user.go" --command="./main"
