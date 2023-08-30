# Server

[rocker/r-ver]() based server with hosted [plumber]() api for R scripts behind nginx reverse proxy with basic auth

Entrypoint is a go CLI that will start plumber after making `plumber.R` file from template based on provided configuration `config.json`

## Environment variables

| NAME        | DESCRIPTION
|-------------|----------------
| SERVER_USER | Nginx basic auth for accessing plumber API
| SERVER_PASS | 

## Mount

| NAME         | DESCRIPTION
|--------------|-----------------
| /config.json | server configuration
| /api/        | directory containing plumber API scripts
| /packages/   | directory containing additional R packages

## Test in docker

```sh
go build github.com/EikenDram/kube-r/server-start
docker build -t kube-r/server:latest .
docker run --rm --name test -p 8000:80 kube-r/server:latest
```

Alright everything works so far