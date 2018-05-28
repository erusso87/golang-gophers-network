FROM golang

# Go Dep & Fresh
RUN go get -u github.com/golang/dep/cmd/dep \
 && go get -u github.com/pilu/fresh

WORKDIR /go/src/gophers-network/gopher-services

ENTRYPOINT ["fresh","-c","../docker/app/dev/runner.conf","main.go"]