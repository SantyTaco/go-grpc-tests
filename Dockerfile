FROM golang:alpine

RUN apk update \
  && apk add git

RUN mkdir -p /usr/goWorkspace/
WORKDIR /usr/goWorkspace/
COPY . .

RUN ls
# Don't do this in production! Use vendoring instead.
# RUN go get -v app/server

#RUN go install app/server

ENTRYPOINT ["go", "run", "greet/greetServer/server.go"]