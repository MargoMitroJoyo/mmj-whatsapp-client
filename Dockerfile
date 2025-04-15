FROM golang:1.24.2 AS builder

WORKDIR /go/src/app

COPY . .

RUN mkdir .store

RUN go mod download

RUN chmod +x ./build.sh

RUN ./build.sh

RUN ldd ./build/bin/main || echo "ldd not available"

FROM gcr.io/distroless/base-debian12 AS final

COPY --from=builder /go/src/app/build/bin/main /app
COPY --from=builder /go/src/app/.store /.store

CMD ["/app"]