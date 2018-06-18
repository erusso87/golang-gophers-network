FROM golang

ADD ./ /go/src/gophers-network
WORKDIR /go/src/gophers-network

# Go Dep
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure \
 && go build -o main

ENTRYPOINT ["./main"]