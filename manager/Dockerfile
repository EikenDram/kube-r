# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /manager

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

# labels
LABEL org.opencontainers.image.source=https://github.com/EikenDram/kube-r
LABEL org.opencontainers.image.description="Container image for running KubeR manager node"
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /

COPY --from=build-stage /manager /manager

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/worker"]