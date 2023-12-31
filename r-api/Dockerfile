# Build the application from source
FROM golang:1.20 AS build-stage

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /r-api-binary

# Deploy the application binary into a r-ver image
FROM rocker/r-ver:4.3.1  AS build-release-stage

# labels
LABEL org.opencontainers.image.source=https://github.com/EikenDram/kube-r
LABEL org.opencontainers.image.description="Container image for running KubeR r-api node"
LABEL org.opencontainers.image.licenses=MIT

# env variables with user and password to access plumber api
ENV SERVER_USER radmin \
    SERVER_PASS radmin

# install the linux libraries needed for plumber
RUN apt-get update -qq && apt-get install -y \
  libssl-dev \
  libcurl4-gnutls-dev

# setup nginx
RUN apt-get update && \
apt-get install -y nginx apache2-utils

# predefined packages
COPY install.R /install.R

# install predefined packages
RUN Rscript install.R

# directory with plumber api scripts
COPY api/ /api/
# directory with additional packages
COPY packages/ /packages/

# templates
COPY templates/ /templates/

# server configuration
COPY config.json /config.json

# plumber api
COPY api.R /api.R
COPY version.R /version.R

# add nginx configuration
ADD ./nginx.conf /etc/nginx/nginx.conf

# server binary
COPY --from=build-stage /r-api-start /r-api-start

# open port 80 to traffic
EXPOSE 80

# when the container starts, start nginx and the server tool
CMD service nginx start && ./r-api-start