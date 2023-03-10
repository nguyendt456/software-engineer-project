FROM golang:1.19

WORKDIR /create_user

#ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
#ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

COPY go.mod go.sum /create_user

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o create_user_service main/create_user/main.go" --command="./create_user_service"
