FROM golang:1.19

WORKDIR /login_user

#ENV GRPC_GO_LOG_VERBOSITY_LEVEL=99
#ENV GRPC_GO_LOG_SEVERITY_LEVEL=info

COPY go.mod go.sum /login_user

RUN go mod download -x

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o login_user_service main/login_user/main.go" --command="./login_user_service"