FROM golang:1.19

WORKDIR /database

#ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
#ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

COPY go.mod go.sum /database

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o database_service main/database/main.go" --command="./database_service"
