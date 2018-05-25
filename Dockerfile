FROM golang

COPY ./ /go/src/gophers-network
WORKDIR /go/src/gophers-network/gopher-services

# Go Dep
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure \
 && go build

ENTRYPOINT ["./gopher-services"]