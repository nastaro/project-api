FROM golang:1.20-alpine AS builder
WORKDIR /dist
COPY . .
RUN go build -o /dist/project-api .
FROM alpine:3.17
WORKDIR /build
COPY --from=builder /dist/project-api /build
RUN addgroup -S martin && adduser -S martin -G martin
USER martin
CMD ["/build/project-api"]