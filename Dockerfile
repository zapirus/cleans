#FROM golang:1.21.1 as builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#
#COPY . .
#
#RUN go build -v -o clean
#
#FROM alpine:latest
#
#COPY --from=builder /app/clean /clean
#
#CMD ["/clean"]
#
#FROM golang:1.21.1 as builder
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#
#COPY . .
#
#RUN go build -v -o clean
#
#FROM alpine:latest
#
#COPY --from=builder /app/clean /app/clean
#WORKDIR /app
#RUN chmod +x /app/clean
#CMD ["/app/clean"]
#FROM golang:1.21.1 as builder
#
#WORKDIR /usr/src/app
#
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify
#
#COPY . .
#
#RUN go build -v ./ && ./clean
#
#
#FROM alpine:latest
#
#COPY --from=builder /usr/src/app/clean /clean
#
#CMD ["/clean"]

FROM golang:1.21.1

RUN go version

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cleans

CMD ["./cleans"]