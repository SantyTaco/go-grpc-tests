FROM golang:alpine

RUN apk update \
  && apk add git

RUN go env
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u google.golang.org/grpc
RUN go get github.com/golang/protobuf/proto
RUN go get firebase.google.com/go
RUN go get google.golang.org/api/option

RUN mkdir -p /usr/goWorkspace/
WORKDIR /usr/goWorkspace/
COPY . .

# Don't do this in production! Use vendoring instead.
# RUN go get -v app/server

#RUN go install app/server

ENTRYPOINT ["go", "run", "/usr/goWorkspace/greetServer/server.go"]