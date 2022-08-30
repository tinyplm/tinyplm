ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

FROM gcr.io/distroless/static-debian11

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 9000

ENTRYPOINT ["./app"]

# docker build -t api-service .
# docker run -d -p 9000:9000 api-service

