FROM golang

COPY ./ /go/src/gophers-network
WORKDIR /go/src/gophers-network/gopher-services

RUN go get github.com/kelseyhightower/envconfig \
 && go get github.com/lib/pq \
 && go get github.com/oklog/ulid \
 && go get github.com/tinrab/retry \
 && go build

ENTRYPOINT ["./gopher-services"]