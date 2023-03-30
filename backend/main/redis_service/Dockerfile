FROM golang:1.19

WORKDIR /redis

#ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
#ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

COPY go.mod go.sum /redis

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o redis_service main/redis_service/main.go" --command="./redis_service"