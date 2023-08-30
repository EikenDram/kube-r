# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.18 AS build-stage

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /worker

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /worker /worker

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/worker"]