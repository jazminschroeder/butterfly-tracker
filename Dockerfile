FROM golang:1.11-alpine as builder
WORKDIR $GOPATH/src/github.com/jazminschroeder/butterflytracker
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o btapp cmd/main.go
RUN ls

FROM scratch
COPY --from=builder go/src/github.com/jazminschroeder/butterflytracker/btapp .
CMD ["./btapp"]