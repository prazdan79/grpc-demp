FROM golang
COPY . /go/src/grpc-demo/client
COPY api /go/src/grpc-demo/client/api
WORKDIR /go/src/grpc-demo/client
RUN go build -i -v -o /go/src/grpc-demo/client .
ENTRYPOINT go run main.go
EXPOSE 8080
