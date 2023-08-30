# KubeR

Service for processing requests for making reports by running R scripts

| COMPONENT | DESCRIPTION
|-----------|-----------------------------------------------------
| web-ui     | Web UI for admin and monitoring
| manager    | manager api for web ui, sending requests into message query, and for calling app api for result status
| worker-api | service for listening to request message query, running server plumber api and calling manager api for status update
| worker-r   | service for listening to request message query, running R scripts and calling manager api for status update
| r-api      | R server hosting plumber api for making reports and api for managing server

Helm chart is available from [https://eikendram.github.io/kube-charts](https://github.com/EikenDram/kube-charts)