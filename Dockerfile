FROM golang:1.21.6-alpine AS builder
RUN  mkdir /build
ADD . /build/
WORKDIR /build
RUN go build ./cmd/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./server"]