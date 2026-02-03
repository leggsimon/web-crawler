# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -ldflags="-w -s" -o /backend .

# Runtime stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /backend /app/backend
EXPOSE 8080
ENTRYPOINT ["/app/backend"]
