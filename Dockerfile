FROM golang:1.24-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /zing

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /zing /zing

EXPOSE 50051

USER nonroot:nonroot

ENTRYPOINT ["/zing", "serve", "-a" , "0.0.0.0", "-p", "50051"]
