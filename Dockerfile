FROM golang:1.24.2 AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN chmod +x ./build.sh

RUN ./build.sh

FROM gcr.io/distroless/base-debian12 AS final

COPY --from=builder /go/src/app/build/bin/main /app

CMD ["/app"]