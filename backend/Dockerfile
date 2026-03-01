FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/server /server
EXPOSE 50051
USER nonroot:nonroot
ENTRYPOINT ["/server"]