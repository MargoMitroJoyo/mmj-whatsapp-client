FROM golang:1.24-alpine AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x ./build.sh

RUN ./build.sh

FROM gcr.io/distroless/static-debian12 AS final

COPY --from=builder /go/src/app/build/bin/main /app

CMD ["/app"]