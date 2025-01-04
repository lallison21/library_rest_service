FROM golang:alpine AS builder

WORKDIR /build

RUN apk --no-cache add bash make
COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

COPY ../.. .

RUN make build

FROM alpine AS runner

COPY --from=builder /build/bin/library_rest_service ./

CMD ["./library_rest_service"]

EXPOSE 8000/tcp
