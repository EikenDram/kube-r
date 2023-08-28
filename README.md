# KubeR

Service for processing requests for making reports by running R scripts

| COMPONENT | DESCRIPTION
|-----------|-----------------------------------------------------
| web-ui    | Web UI for admin and monitoring
| server    | R server hosting plumber api for making reports and api for managing server
| manager   | manager api for web ui, sending requests into message query, and for calling app api for result status
| worker    | service for listening to request message query, running server plumber api and calling api in manager for result status

Helm chart is available from [https://eikendram.github.io/kube-charts](https://github.com/EikenDram/kube-charts)

## Web-UI

## Server

## Manager

## Worker

Q: Security - either user-password or add keycloak? think i'll go with user-password for now and then optional OAuth