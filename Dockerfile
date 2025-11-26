# ---------------------------------------------------------
# Build stage
# ---------------------------------------------------------
FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /zing

# ---------------------------------------------------------
# Final runtime image
# ---------------------------------------------------------
FROM alpine:latest AS build-release-stage

# Create a non-root user and group
RUN addgroup -S app && adduser -S app -G app

WORKDIR /

# Copy binary from builder
COPY --from=build-stage /zing /zing

EXPOSE 50051

# Switch to non-root user
USER app:app

ENTRYPOINT ["/zing", "serve", "-a" , "0.0.0.0", "-p", "50051"]
