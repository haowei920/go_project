FROM golang:1.16-buster

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/haowei920/go_project/server
RUN cd /build && git clone https://github.com/haowei920/go_project.git

RUN cd /build/go_project/client && chmod +x cli.sh
RUN cd /build/go_project/server && chmod +x server.sh
RUN cd /build/go_project/server && go build server.go

# EXPOSE 8080

ENTRYPOINT ["/build/go_project/server/server"]


